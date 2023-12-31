// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
package main

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/containrrr/shoutrrr"
	"github.com/containrrr/shoutrrr/pkg/router"
	delugeclient "github.com/gdm85/go-libdeluge"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lrstanley/clix"
	"github.com/lrstanley/deluge-remove-after/internal/models"
	"github.com/lrstanley/deluge-remove-after/internal/worker"
)

var (
	version = "master"
	commit  = "latest"
	date    = "-"

	logger log.Interface
	cli    = &clix.CLI[models.Flags]{
		Links: clix.GithubLinks("github.com/lrstanley/deluge-remove-after", "master", "https://liam.sh"),
		VersionInfo: &clix.VersionInfo[models.Flags]{
			Name:    "deluge-remove-after",
			Version: version,
			Commit:  commit,
			Date:    date,
		},
	}

	notifiers *router.ServiceRouter
)

func main() {
	cli.Parse()
	logger = cli.Logger

	ctx := log.NewContext(context.Background(), logger)
	var err error

	if cli.Flags.Deluge.CheckInterval < 5*time.Minute {
		cli.Flags.Deluge.CheckInterval = 5 * time.Minute
	}

	if cli.Flags.Deluge.MaxSeedTime > 0 && cli.Flags.Deluge.MaxSeedTime < 1*time.Hour {
		cli.Flags.Deluge.MaxSeedTime = 1 * time.Hour
	}

	if cli.Flags.Deluge.MaxTimeAdded > 0 && cli.Flags.Deluge.MaxTimeAdded < 6*time.Hour {
		cli.Flags.Deluge.MaxTimeAdded = 6 * time.Hour
	}

	if len(cli.Flags.Notifiers) > 0 {
		notifiers, err = shoutrrr.CreateSender(cli.Flags.Notifiers...)
		if err != nil {
			logger.WithError(err).Fatal("unable to parse notifiers")
		}
	}

	deluge := delugeclient.NewV2(delugeclient.Settings{
		Hostname:         cli.Flags.Deluge.Hostname,
		Port:             cli.Flags.Deluge.Port,
		Login:            cli.Flags.Deluge.Username,
		Password:         cli.Flags.Deluge.Password,
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

	if cli.Flags.Deluge.RemoveTorrent {
		states = append(states, delugeclient.StatePaused)
	}

	w := &worker.Worker{
		DryRun:    cli.Flags.DryRun,
		Config:    cli.Flags.Deluge,
		Client:    deluge,
		States:    states,
		Notifiers: notifiers,
	}

	if cli.Flags.NoDaemon {
		err = w.CheckStates(ctx)
		if err != nil {
			logger.WithError(err).Fatal("error running worker")
		}
		return
	}

	err = clix.RunCtx(ctx, w.Watcher)
	if err != nil {
		logger.WithError(err).Fatal("error running worker")
	}
}
