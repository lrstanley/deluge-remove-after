<!-- template:begin:header -->
<!-- do not edit anything in this "template" block, its auto-generated -->
<p align="center">deluge-remove-after -- daemon for removing torrents from deluge after a specified timeframe</p>
<p align="center">




  <a href="https://github.com/lrstanley/deluge-remove-after/actions?query=workflow%3Atest+event%3Apush">
    <img alt="GitHub Workflow Status (test @ master)" src="https://img.shields.io/github/workflow/status/lrstanley/deluge-remove-after/test/master?label=test&style=flat-square&event=push">
  </a>

  <img alt="Code Coverage" src="https://img.shields.io/codecov/c/github/lrstanley/deluge-remove-after/master?style=flat-square">

  <a href="https://pkg.go.dev/github.com/lrstanley/deluge-remove-after">
    <img alt="Go Documentation" src="https://pkg.go.dev/badge/github.com/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/lrstanley/deluge-remove-after">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <img alt="Bug reports" src="https://img.shields.io/github/issues/lrstanley/deluge-remove-after/bug?label=issues&style=flat-square">
  <img alt="Feature requests" src="https://img.shields.io/github/issues/lrstanley/deluge-remove-after/enhancement?label=feature%20requests&style=flat-square">
  <a href="https://github.com/lrstanley/deluge-remove-after/pulls">
    <img alt="Open Pull Requests" src="https://img.shields.io/github/issues-pr/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/deluge-remove-after/tags">
    <img alt="Latest Semver Tag" src="https://img.shields.io/github/v/tag/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <img alt="Last commit" src="https://img.shields.io/github/last-commit/lrstanley/deluge-remove-after?style=flat-square">
  <a href="https://github.com/lrstanley/deluge-remove-after/discussions/new?category=q-a">
    <img alt="Ask a Question" src="https://img.shields.io/badge/discussions-ask_a_question!-green?style=flat-square">
  </a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/discord-bytecord-blue.svg?style=flat-square" alt="Discord Chat"></a>
</p>
<!-- template:end:header -->

<!-- template:begin:toc -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :link: Table of Contents

  - [Why](#grey_question-why)
  - [Installation](#computer-installation)
    - [Source](#toolbox-source)
  - [Usage](#gear-usage)
  - [License](#balance_scale-license)
<!-- template:end:toc -->

## :grey_question: Why

Deluge only supports a handful of options for pausing torrents after a specific
seed ratio, but not based off time. This should allow you to combine the built-in
seed ratio removal, with the ability to also remove after a certain amount of
time seeded.

## :computer: Installation

Check out the [releases](https://github.com/users/lrstanley/deluge-remove-after/pkgs/container/deluge-remove-after)
page for prebuilt versions.

<!-- template:begin:ghcr -->
<!-- do not edit anything in this "template" block, its auto-generated -->

<!-- template:end:ghcr -->

### :toolbox: Source

Note that you must have [Go](https://golang.org/doc/install) installed (latest is usually best).

    $ git clone https://github.com/lrstanley/deluge-remove-after.git && cd deluge-remove-after
    $ make
    $ ./deluge-remove-after --help

## :gear: Usage

Take a look at the [docker-compose.yaml](/docker-compose.yaml) file, or the above
docker run commands. For references on supported flags/environment variables,
take a look at [USAGE.md](/USAGE.md)

<!-- template:begin:support -->
<!-- template:end:support -->

<!-- template:begin:contributing -->
<!-- template:end:contributing -->

<!-- template:begin:license -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :balance_scale: License

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

_Also located [here](LICENSE)_
<!-- template:end:license -->
