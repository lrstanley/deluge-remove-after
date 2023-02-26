// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package worker

import (
	"context"
	"fmt"
	"time"

	"github.com/apex/log"
	"github.com/containrrr/shoutrrr/pkg/router"
	"github.com/containrrr/shoutrrr/pkg/types"
	delugeclient "github.com/gdm85/go-libdeluge"
	"github.com/lrstanley/deluge-remove-after/internal/models"
)

type Worker struct {
	DryRun bool
	Config models.ConfigDeluge

	Client    *delugeclient.ClientV2
	States    []delugeclient.TorrentState
	Notifiers *router.ServiceRouter
}

func (w *Worker) Watcher(ctx context.Context) error {
	ticker := time.NewTicker(w.Config.CheckInterval)
	defer ticker.Stop()

	// Run once on startup.
	if err := w.checkStates(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if err := w.checkStates(ctx); err != nil {
				return err
			}
		}
	}
}

func (w *Worker) checkStates(ctx context.Context) error {
	l := log.FromContext(ctx)

	l.Info("running checks")
	defer l.Info("checks completed")

	for _, state := range w.States {
		logEntry := l.WithField("state", state)

		logEntry.Info("worker running")

		err := w.process(log.NewContext(ctx, logEntry), state)
		if err != nil {
			logEntry.WithError(err).Error("worker processing failed")

			if w.Notifiers != nil {
				w.Notifiers.Send(
					fmt.Sprintf("error processing torrents for state[%s]: %v", state, err),
					&types.Params{},
				)
			}
		} else {
			logEntry.Info("worker completed")
		}
	}

	return nil
}

func (w *Worker) process(ctx context.Context, state delugeclient.TorrentState) error {
	out, err := w.Client.TorrentsStatus(state, nil)
	if err != nil {
		return err
	}

	l := log.FromContext(ctx)

	changed := 0

	for hash := range out {
		if !out[hash].IsFinished {
			l.WithFields(log.Fields{
				"hash": hash,
				"name": out[hash].Name,
			}).Debug("skipping unfinished torrent")

			continue
		}

		durationSeeding := time.Duration(out[hash].SeedingTime) * time.Second
		durationSinceAdded := time.Since(time.Unix(int64(out[hash].TimeAdded), 0))

		logEntry := l.WithFields(log.Fields{
			"hash": hash,
			"name": out[hash].Name,
		})

		if w.Config.MaxSeedTime > 0 && durationSeeding > w.Config.MaxSeedTime {
			logEntry.WithField("seeding", durationSeeding).Debug("hit max seed time")

			changed++
			w.removeTorrent(log.NewContext(ctx, logEntry), hash, out[hash])
			continue
		}

		if w.Config.MaxTimeAdded > 0 && durationSinceAdded > w.Config.MaxTimeAdded {
			logEntry.WithField("time_since_added", durationSinceAdded).Debug("hit max time added")

			changed++
			w.removeTorrent(log.NewContext(ctx, logEntry), hash, out[hash])
			continue
		}
	}

	l.Infof("%d of %d torrents changed and/or matched criteria for given state", changed, len(out))
	return nil
}

func (w *Worker) removeTorrent(ctx context.Context, hash string, torrent *delugeclient.TorrentStatus) {
	if w.DryRun {
		return
	}

	l := log.FromContext(ctx)

	if w.Config.RemoveTorrent {
		_, err := w.Client.RemoveTorrent(hash, w.Config.RemoveFiles)
		if err != nil {
			l.WithError(err).Error("unable to remove torrent")
			return
		}
		l.WithField("rm_files", w.Config.RemoveFiles).Info("removed torrent")
		return
	}

	err := w.Client.PauseTorrents(hash)
	if err != nil {
		l.WithError(err).Error("unable to pause torrent")
	}
	l.Info("paused torrent")
}
