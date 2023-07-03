<!-- template:begin:header -->
<!-- do not edit anything in this "template" block, its auto-generated -->

<p align="center">deluge-remove-after -- daemon for removing torrents from deluge after a specified timeframe</p>
<p align="center">
  <a href="https://github.com/lrstanley/deluge-remove-after/tags">
    <img title="Latest Semver Tag" src="https://img.shields.io/github/v/tag/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/deluge-remove-after/commits/master">
    <img title="Last commit" src="https://img.shields.io/github/last-commit/lrstanley/deluge-remove-after?style=flat-square">
  </a>




  <a href="https://github.com/lrstanley/deluge-remove-after/actions?query=workflow%3Atest+event%3Apush">
    <img title="GitHub Workflow Status (test @ master)" src="https://img.shields.io/github/actions/workflow/status/lrstanley/deluge-remove-after/test.yml?branch=master&label=test&style=flat-square">
  </a>

  <a href="https://codecov.io/gh/lrstanley/deluge-remove-after">
    <img title="Code Coverage" src="https://img.shields.io/codecov/c/github/lrstanley/deluge-remove-after/master?style=flat-square">
  </a>

  <a href="https://pkg.go.dev/github.com/lrstanley/deluge-remove-after">
    <img title="Go Documentation" src="https://pkg.go.dev/badge/github.com/lrstanley/deluge-remove-after?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/lrstanley/deluge-remove-after">
    <img title="Go Report Card" src="https://goreportcard.com/badge/github.com/lrstanley/deluge-remove-after?style=flat-square">
  </a>
</p>
<p align="center">
  <a href="https://github.com/lrstanley/deluge-remove-after/issues?q=is:open+is:issue+label:bug">
    <img title="Bug reports" src="https://img.shields.io/github/issues/lrstanley/deluge-remove-after/bug?label=issues&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/deluge-remove-after/issues?q=is:open+is:issue+label:enhancement">
    <img title="Feature requests" src="https://img.shields.io/github/issues/lrstanley/deluge-remove-after/enhancement?label=feature%20requests&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/deluge-remove-after/pulls">
    <img title="Open Pull Requests" src="https://img.shields.io/github/issues-pr/lrstanley/deluge-remove-after?label=prs&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/deluge-remove-after/discussions/new?category=q-a">
    <img title="Ask a Question" src="https://img.shields.io/badge/support-ask_a_question!-blue?style=flat-square">
  </a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/discord-bytecord-blue.svg?style=flat-square" title="Discord Chat"></a>
</p>
<!-- template:end:header -->

<!-- template:begin:toc -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :link: Table of Contents

  - [Why](#grey_question-why)
  - [Similar projects](#raising_hand_man-similar-projects)
  - [💻 Installation](#computer-installation)
    - [🐳 Container Images (ghcr)](#whale-container-images-ghcr)
    - [🧰 Source](#toolbox-source)
  - [Usage](#gear-usage)
  - [Support &amp; Assistance](#raising_hand_man-support--assistance)
  - [Contributing](#handshake-contributing)
  - [⚖️ License](#balance_scale-license)
<!-- template:end:toc -->

## :grey_question: Why

Deluge only supports a handful of options for pausing torrents after a specific
seed ratio, but not based off time. This should allow you to combine the built-in
seed ratio removal, with the ability to also remove after a certain amount of
time seeded.

## :raising_hand_man: Similar projects

- [laur89/deluge-autoremoveplus](https://github.com/laur89/deluge-autoremoveplus) -- patched
  version of the orignal AutoRemovePlus for Deluge

## :computer: Installation

Check out the [releases](https://github.com/users/lrstanley/deluge-remove-after/pkgs/container/deluge-remove-after)
page for prebuilt versions.

<!-- template:begin:ghcr -->
<!-- do not edit anything in this "template" block, its auto-generated -->
### :whale: Container Images (ghcr)

```console
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:master
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:2.0.0
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:latest
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:1.0.1
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:1.0.0
$ docker run -it --rm ghcr.io/lrstanley/deluge-remove-after:0.0.1
```
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
<!-- do not edit anything in this "template" block, its auto-generated -->
## :raising_hand_man: Support & Assistance

* :heart: Please review the [Code of Conduct](.github/CODE_OF_CONDUCT.md) for
     guidelines on ensuring everyone has the best experience interacting with
     the community.
* :raising_hand_man: Take a look at the [support](.github/SUPPORT.md) document on
     guidelines for tips on how to ask the right questions.
* :lady_beetle: For all features/bugs/issues/questions/etc, [head over here](https://github.com/lrstanley/deluge-remove-after/issues/new/choose).
<!-- template:end:support -->

<!-- template:begin:contributing -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :handshake: Contributing

* :heart: Please review the [Code of Conduct](.github/CODE_OF_CONDUCT.md) for guidelines
     on ensuring everyone has the best experience interacting with the
    community.
* :clipboard: Please review the [contributing](.github/CONTRIBUTING.md) doc for submitting
     issues/a guide on submitting pull requests and helping out.
* :old_key: For anything security related, please review this repositories [security policy](https://github.com/lrstanley/deluge-remove-after/security/policy).
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
