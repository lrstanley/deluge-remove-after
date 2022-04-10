<!-- template:begin:header -->
<!-- template:end:header -->

<!-- template:begin:toc -->
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
<!-- template:end:license -->
