package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/discard"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/containrrr/shoutrrr/pkg/types"
	delugeclient "github.com/gdm85/go-libdeluge"
	flags "github.com/jessevdk/go-flags"
	_ "github.com/joho/godotenv/autoload"
)

var (
	version = "master"
	commit  = "latest"
	date    = "-"

	cli    = &Flags{}
	logger = &log.Logger{}

	notifiers *router.ServiceRouter
)

type Flags struct {
	VersionFlag bool `short:"v" long:"version" description:"display the version and exit"`
	Debug       bool `env:"DEBUG" short:"D" long:"debug" description:"enable bot debugging"`

	DryRun bool `env:"DRY_RUN" long:"dry-run" description:"dry-run operations (does NOT change/remove torrents)"`

	Notifiers []string `env:"NOTIFIERS" long:"notifiers" description:"list of shoutrrr notification urls: https://containrrr.dev/shoutrrr/)"`

	Deluge struct {
		Username string `env:"USERNAME" short:"u" long:"username" default:"localclient" description:"deluge username (NOT web-ui username)"`
		Password string `env:"PASSWORD" short:"p" long:"password" description:"deluge password (NOT web-ui password)"`

		Hostname string `env:"HOSTNAME" short:"H" long:"hostname" default:"localhost" description:"deluge hostname"`
		Port     uint   `env:"PORT" short:"P" long:"port" default:"58846" description:"deluge port"`

		RemoveTorrent bool          `env:"REMOVE_TORRENT" short:"r" long:"remove-torrent" description:"Remove torrent (default pauses torrent)"`
		RemoveFiles   bool          `env:"REMOVE_FILES" short:"R" long:"remove-files" description:"if true, when removing a torrent (see: --remove-torrent), the torrent files will be removed as well"`
		CheckInterval time.Duration `env:"CHECK_INTERVAL" short:"i" long:"check-interval" default:"6h" description:"how often to check torrent statuses (format: s, m h)"`

		MaxSeedTime  time.Duration `env:"MAX_SEED_TIME" short:"S" long:"max-seed-time" default:"336h" description:"max time a completed torrent can be seeded for (format: s, m h)"`
		MaxTimeAdded time.Duration `env:"MAX_TIME_ADDED" short:"M" long:"max-time-added" description:"amount of time after a completed torrent was added to deluge, before it should be removed (format: s, m h)"`
	} `group:"Deluge & Torrent Options" namespace:"deluge" env-namespace:"DELUGE"`

	Log struct {
		Quiet bool   `env:"LOG_QUIET" long:"quiet" description:"disable logging to stdout (also: see levels)"`
		Level string `env:"LOG_LEVEL" long:"level" default:"info" choice:"debug" choice:"info" choice:"warn" choice:"error" choice:"fatal"  description:"logging level"`
		JSON  bool   `env:"LOG_JSON" long:"json" description:"output logs in JSON format"`
	} `group:"Logging Options" namespace:"log" env-namespace:"LOG"`
}

func main() {
	var err error

	parser := flags.NewParser(cli, flags.HelpFlag|flags.PrintErrors|flags.PassDoubleDash)
	parser.NamespaceDelimiter = "."
	parser.EnvNamespaceDelimiter = "_"

	if _, err = parser.Parse(); err != nil {
		// fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if cli.VersionFlag {
		fmt.Printf("deluge-remove-after version: %s [%s] (%s, %s), compiled %s\n", version, commit, runtime.GOOS, runtime.GOARCH, date)
		os.Exit(1)
	}

	// Initialize logging.
	if cli.Debug {
		logger.Level = log.DebugLevel
	} else {
		logger.Level = log.MustParseLevel(cli.Log.Level)
	}

	if cli.Log.Quiet {
		logger.Handler = discard.New()
	} else if cli.Log.JSON {
		logger.Handler = json.New(os.Stdout)
	} else {
		logger.Handler = text.New(os.Stdout)
	}

	if cli.Deluge.CheckInterval < 5*time.Minute {
		cli.Deluge.CheckInterval = 5 * time.Minute
	}

	if cli.Deluge.MaxSeedTime > 0 && cli.Deluge.MaxSeedTime < 1*time.Hour {
		cli.Deluge.MaxSeedTime = 1 * time.Hour
	}

	if cli.Deluge.MaxTimeAdded > 0 && cli.Deluge.MaxTimeAdded < 6*time.Hour {
		cli.Deluge.MaxTimeAdded = 6 * time.Hour
	}

	if len(cli.Notifiers) > 0 {
		notifiers, err = shoutrrr.CreateSender(cli.Notifiers...)
		if err != nil {
			logger.WithError(err).Fatal("unable to parse notifiers")
		}
	}

	deluge := delugeclient.NewV2(delugeclient.Settings{
		Hostname:         cli.Deluge.Hostname,
		Port:             cli.Deluge.Port,
		Login:            cli.Deluge.Username,
		Password:         cli.Deluge.Password,
		ReadWriteTimeout: 30 * time.Second,
	})

	// perform connection to Deluge server
	err = deluge.Connect()
	if err != nil {
		logger.WithError(err).Fatal("unable to connect to deluge daemon")
	}

	states := []delugeclient.TorrentState{
		delugeclient.StateActive,
		delugeclient.StateSeeding,
		delugeclient.StateQueued,
	}

	if cli.Deluge.RemoveTorrent {
		states = append(states, delugeclient.StatePaused)
	}

	for {
		logger.Info("running checks")
		defer logger.Info("checks completed")

		for _, state := range states {
			logger.WithField("state", state).Info("worker running")
			err = process(deluge, state)
			if err != nil {
				logger.WithField("state", state).WithError(err).Error("worker processing failed")

				if notifiers != nil {
					notifiers.Send(
						fmt.Sprintf("error processing torrents for state[%s]: %v", state, err),
						&types.Params{},
					)
				}
			} else {
				logger.WithField("state", state).Info("worker completed")
			}
		}

		logger.WithField("duration", cli.Deluge.CheckInterval).Info("sleeping for duration")
		time.Sleep(cli.Deluge.CheckInterval)
	}
}

func process(client *delugeclient.ClientV2, state delugeclient.TorrentState) error {
	out, err := client.TorrentsStatus(state, nil)
	if err != nil {
		return err
	}

	baseEntry := logger.WithFields(log.Fields{
		"state": state,
	})

	changed := 0

	for hash := range out {
		if !out[hash].IsFinished {
			logger.WithFields(log.Fields{
				"state": state,
				"hash":  hash,
				"name":  out[hash].Name,
			}).Debug("skipping unfinished torrent")

			continue
		}

		durationSeeding := time.Duration(out[hash].SeedingTime) * time.Second
		durationSinceAdded := time.Since(time.Unix(int64(out[hash].TimeAdded), 0))

		logEntry := baseEntry.WithFields(log.Fields{
			"hash": hash,
			"name": out[hash].Name,
		})

		if cli.Deluge.MaxSeedTime > 0 && durationSeeding > cli.Deluge.MaxSeedTime {
			logEntry.WithField("seeding", durationSeeding).Debug("hit max seed time")

			changed++
			removeTorrent(client, hash, out[hash])
			continue
		}

		if cli.Deluge.MaxTimeAdded > 0 && durationSinceAdded > cli.Deluge.MaxTimeAdded {
			logEntry.WithField("time_since_added", durationSinceAdded).Debug("hit max time added")

			changed++
			removeTorrent(client, hash, out[hash])
			continue
		}
	}

	baseEntry.Infof("%d of %d torrents changed and/or matched criteria for given state", changed, len(out))
	return nil
}

func removeTorrent(client *delugeclient.ClientV2, hash string, torrent *delugeclient.TorrentStatus) {
	if cli.DryRun {
		return
	}

	logEntry := logger.WithFields(log.Fields{
		"hash": hash,
		"name": torrent.Name,
	})

	if cli.Deluge.RemoveTorrent {
		_, err := client.RemoveTorrent(hash, cli.Deluge.RemoveFiles)
		if err != nil {
			logEntry.WithError(err).Error("unable to remove torrent")
			return
		}
		logEntry.WithField("rm_files", cli.Deluge.RemoveFiles).Info("removed torrent")
		return
	}

	err := client.PauseTorrents(hash)
	if err != nil {
		logEntry.WithError(err).Error("unable to pause torrent")
	}
	logEntry.Info("paused torrent")
}
