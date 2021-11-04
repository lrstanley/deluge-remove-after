<p align="center">deluge-remove-after -- daemon for removing torrents from deluge after a specified timeframe</p>
<p align="center">
  <a href="https://github.com/lrstanley/deluge-remove-after/releases"><img src="https://github.com/lrstanley/deluge-remove-after/workflows/release/badge.svg" alt="Release Status"></a>
  <a href="https://github.com/users/lrstanley/packages/container/deluge-remove-after/versions"><img src="https://img.shields.io/badge/Docker-ghcr.io%2Flrstanley%2Fdeluge--remove--after-blue.svg" alt="Docker"></a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/Community-Chat%20with%20us-green.svg" alt="Community Chat"></a>
</p>

## Table of Contents
- [Why](#why)
- [Installation](#installation)
  - [Docker](#docker)
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
page for prebuilt versions. Below are example commands of how you would install
the utility.

### Docker

```bash
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:latest
```

### Source

Note that you must have [Go](https://golang.org/doc/install) installed (latest is usually best).

    $ git clone https://github.com/lrstanley/deluge-remove-after.git && cd deluge-remove-after
    $ make
    $ ./deluge-remove-after --help

## Usage

Take a look at the [docker-compose.yaml] file.

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
