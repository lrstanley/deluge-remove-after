## :gear: Usage

#### Application Options
| Environment vars | Flags | Type | Description |
| --- | --- | --- | --- |
| `DRY_RUN` | `--dry-run` | bool | dry-run operations (does NOT change/remove torrents) |
| `NOTIFIERS` | `--notifiers` | []string | list of shoutrrr notification urls: https://containrrr.dev/shoutrrr/) |
| `NO_DAEMON` | `--no-daemon` | bool | do not run as a daemon (run once and exit) |
| - | `-v, --version` | bool | prints version information and exits |
| - | `--version-json` | bool | prints version information in JSON format and exits |
| `DEBUG` | `-D, --debug` | bool | enables debug mode |

#### Deluge & Torrent Options
| Environment vars | Flags | Type | Description |
| --- | --- | --- | --- |
| `DELUGE_USERNAME` | `-u, --deluge.username` | string | deluge username (NOT web-ui username) [**default: localclient**] |
| `DELUGE_PASSWORD` | `-p, --deluge.password` | string | deluge password (NOT web-ui password) |
| `DELUGE_HOSTNAME` | `-H, --deluge.hostname` | string | deluge hostname [**default: localhost**] |
| `DELUGE_PORT` | `-P, --deluge.port` | uint | deluge port [**default: 58846**] |
| `DELUGE_REMOVE_TORRENT` | `-r, --deluge.remove-torrent` | bool | Remove torrent (default pauses torrent) |
| `DELUGE_REMOVE_FILES` | `-R, --deluge.remove-files` | bool | if true, when removing a torrent (see: --remove-torrent), the torrent files will be removed as well |
| `DELUGE_CHECK_INTERVAL` | `-i, --deluge.check-interval` | time.Duration | how often to check torrent statuses (format: s, m h) [**default: 6h**] |
| `DELUGE_MAX_SEED_TIME` | `-S, --deluge.max-seed-time` | time.Duration | max time a completed torrent can be seeded for (format: s, m h) [**default: 336h**] |
| `DELUGE_MAX_TIME_ADDED` | `-M, --deluge.max-time-added` | time.Duration | amount of time after a completed torrent was added to deluge, before it should be removed (format: s, m h) |

#### Logging Options
| Environment vars | Flags | Type | Description |
| --- | --- | --- | --- |
| `LOG_QUIET` | `--log.quiet` | bool | disable logging to stdout (also: see levels) |
| `LOG_LEVEL` | `--log.level` | string | logging level [**default: info**] [**choices: debug, info, warn, error, fatal**] |
| `LOG_JSON` | `--log.json` | bool | output logs in JSON format |
| `LOG_PRETTY` | `--log.pretty` | bool | output logs in a pretty colored format (cannot be easily parsed) |
| `LOG_PATH` | `--log.path` | string | path to log file (disables stdout logging) |
