// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package models

import "time"

type Flags struct {
	DryRun    bool     `env:"DRY_RUN" long:"dry-run" description:"dry-run operations (does NOT change/remove torrents)"`
	Notifiers []string `env:"NOTIFIERS" long:"notifiers" description:"list of shoutrrr notification urls: https://containrrr.dev/shoutrrr/)"`

	Deluge ConfigDeluge `group:"Deluge & Torrent Options" namespace:"deluge" env-namespace:"DELUGE"`
}

type ConfigDeluge struct {
	Username string `env:"USERNAME" short:"u" long:"username" default:"localclient" description:"deluge username (NOT web-ui username)"`
	Password string `env:"PASSWORD" short:"p" long:"password" description:"deluge password (NOT web-ui password)"`
	Hostname string `env:"HOSTNAME" short:"H" long:"hostname" default:"localhost" description:"deluge hostname"`
	Port     uint   `env:"PORT"     short:"P" long:"port"     default:"58846" description:"deluge port"`

	RemoveTorrent bool          `env:"REMOVE_TORRENT" short:"r" long:"remove-torrent" description:"Remove torrent (default pauses torrent)"`
	RemoveFiles   bool          `env:"REMOVE_FILES"   short:"R" long:"remove-files"   description:"if true, when removing a torrent (see: --remove-torrent), the torrent files will be removed as well"`
	CheckInterval time.Duration `env:"CHECK_INTERVAL" short:"i" long:"check-interval" default:"6h" description:"how often to check torrent statuses (format: s, m h)"`
	MaxSeedTime   time.Duration `env:"MAX_SEED_TIME"  short:"S" long:"max-seed-time"  default:"336h" description:"max time a completed torrent can be seeded for (format: s, m h)"`
	MaxTimeAdded  time.Duration `env:"MAX_TIME_ADDED" short:"M" long:"max-time-added" description:"amount of time after a completed torrent was added to deluge, before it should be removed (format: s, m h)"`
}
