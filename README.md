<p align="center">deluge-remove-after -- remove torrents from deluge after a specified timeframe</p>
<p align="center">
  <a href="https://github.com/lrstanley/deluge-remove-after/releases"><img src="https://github.com/lrstanley/deluge-remove-after/workflows/release/badge.svg" alt="Release Status"></a>
  <a href="https://github.com/lrstanley/deluge-remove-after/actions"><img src="https://github.com/lrstanley/deluge-remove-after/workflows/build/badge.svg" alt="Build Status"></a>
  <a href="https://github.com/users/lrstanley/packages/container/deluge-remove-after/versions"><img src="https://img.shields.io/badge/Docker-lrstanley%2Fdeluge--remove--after-blue.svg" alt="Docker"></a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/Community-Chat%20with%20us-green.svg" alt="Community Chat"></a>
</p>

## Table of Contents
- [Why](#why)
- [Installation](#installation)
  - [Docker](#docker)
  - [Ubuntu/Debian](#ubuntudebian)
  - [CentOS/Redhat](#centosredhat)
  - [Manual Install](#manual-install)
  - [Build from source](#build-from-source)
- [Usage](#usage)
- [Contributing](#contributing)
- [TODO](#todo)
- [License](#license)

## Why

Deluge only supports a handful of options for pausing torrents after a specific
seed ratio, but not based off time. This should allow you to combine the built-in
seed ratio removal, with the ability to also remove after a certain amount of
time seeded.

## Installation

Check out the [releases](https://github.com/users/lrstanley/deluge-remove-after/pkgs/container/deluge-remove-after)
page for prebuilt versions. deluge-remove-after should work on ubuntu/debian,
centos/redhat/fedora, etc. Below are example commands of how you would install
the utility.

### Docker

```bash
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:latest
```

### Ubuntu/Debian

```bash
$ wget https://liam.sh/ghr/deluge-remove-after_<version>_linux_amd64.deb
$ dpkg -i deluge-remove-after_<version>_linux_amd64.deb
$ deluge-remove-after --help
```

### CentOS/Redhat

```bash
$ yum localinstall https://liam.sh/ghr/deluge-remove-after_<version>_linux_amd64.rpm
$ deluge-remove-after --help
```

Some older CentOS versions may require (if you get `Cannot open: <url>. Skipping.`):

```console
$ wget https://liam.sh/ghr/deluge-remove-after_<version>_linux_amd64.rpm
$ yum localinstall deluge-remove-after_<version>_linux_amd64.rpm
```

### Manual Install

```bash
$ wget https://liam.sh/ghr/deluge-remove-after_<version>_linux_amd64.tar.gz
$ tar -C /usr/bin/ -xzvf deluge-remove-after_<version>_linux_amd64.tar.gz deluge-remove-after
$ chmod +x /usr/bin/deluge-remove-after
$ deluge-remove-after --help
```

### Source

Note that you must have [Go](https://golang.org/doc/install) installed (latest is usually best).

    $ git clone https://github.com/lrstanley/deluge-remove-after.git && cd deluge-remove-after
    $ make
    $ ./deluge-remove-after --help

## Usage

Take a look at the [docker-compose.yaml] file. Note, all fields can be provided
via environment variables (deluge-remove-after also supports `.env` files).

```
$ deluge-remove-after --help
Usage:
  deluge-remove-after [OPTIONS]

Application Options:
  -v, --version                                 display the version and exit
  -D, --debug                                   enable bot debugging [$DEBUG]
      --dry-run                                 dry-run operations (does NOT change/remove torrents)
                                                [$DRY_RUN]

Deluge & Torrent Options:
  -u, --deluge.username=                        deluge username (NOT web-ui username) (default: localclient)
                                                [$DELUGE_USERNAME]
  -p, --deluge.password=                        deluge password (NOT web-ui password) [$DELUGE_PASSWORD]
  -H, --deluge.hostname=                        deluge hostname (default: localhost) [$DELUGE_HOSTNAME]
  -P, --deluge.port=                            deluge port (default: 58846) [$DELUGE_PORT]
  -r, --deluge.remove-torrent                   Remove torrent (default pauses torrent)
                                                [$DELUGE_REMOVE_TORRENT]
  -R, --deluge.remove-files                     if true, when removing a torrent (see: --remove-torrent),
                                                the torrent files will be removed as well
                                                [$DELUGE_REMOVE_FILES]
  -i, --deluge.check-interval=                  how often to check torrent statuses (format: s, m h)
                                                (default: 6h) [$DELUGE_CHECK_INTERVAL]
  -S, --deluge.max-seed-time=                   max time a completed torrent can be seeded for (format: s, m
                                                h) (default: 336h) [$DELUGE_MAX_SEED_TIME]
  -M, --deluge.max-time-added=                  amount of time after a completed torrent was added to
                                                deluge, before it should be removed (format: s, m h)
                                                [$DELUGE_MAX_TIME_ADDED]

Logging Options:
      --log.quiet                               disable logging to stdout (also: see levels) [$LOG_LOG_QUIET]
      --log.level=[debug|info|warn|error|fatal] logging level (default: info) [$LOG_LOG_LEVEL]
      --log.json                                output logs in JSON format [$LOG_LOG_JSON]

Help Options:
  -h, --help                                    Show this help message
```

## Contributing

Please review the [CONTRIBUTING](CONTRIBUTING.md) doc for submitting issues/a guide
on submitting pull requests and helping out.


## License

```
MIT License

Copyright (c) 2021 Liam Stanley <me@liamstanley.io>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
