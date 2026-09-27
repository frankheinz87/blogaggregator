# blogaggregator

A command-line RSS feed aggregator written in Go. `blogaggregator` stores feeds and posts in PostgreSQL, periodically fetches feed updates, and lets registered users follow feeds and browse their posts.

## Prerequisites

- **Go 1.25.3 or newer** (the version declared by this project is Go 1.25.3).
- **PostgreSQL**, with a server running and a database created for the application.
- **Goose**, to apply the database migrations. Goose is a separate tool and is not bundled with the CLI.

## Install

Install the latest published version with Go:

```sh
go install github.com/frankheinz87/blogaggregator@latest
```

Go installs the executable as `blogaggregator`. Make sure Go's binary directory is on your `PATH`. By default, this is `$(go env GOPATH)/bin` (unless `GOBIN` is set).

To use the shorter `gator` command in your current shell, define an alias:

```sh
alias gator=blogaggregator
```

Add that line to your shell's startup file (for example, `~/.bashrc` or `~/.zshrc`) to keep the alias for future terminal sessions. The examples below use `gator`; use `blogaggregator` instead if you do not set up the alias.

## Set Up PostgreSQL

Create a PostgreSQL database for the application. For example, if PostgreSQL is running locally and its command-line tools are available:

```sh
createdb gator
```

Install Goose:

```sh
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Apply the project's migrations from the repository root, substituting your PostgreSQL username and password:

```sh
goose -dir ./sql/schema postgres "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable" up
```

The database in the connection URL must already exist. If your username or password contains URL-reserved characters, URL-encode them in the connection string.

## Configure

The CLI reads its database connection URL from `~/.gatorconfig.json`. Create that file with your own PostgreSQL credentials and database name:

```json
{
	"db_url": "postgres://USER:PASSWORD@localhost:5432/gator?sslmode=disable",
	"current_user_name": ""
}
```

The program updates `current_user_name` when you register or log in. Keep this file in your home directory; the application expects it there when any command starts.

## Quick Start

Register a user. Registration also makes that user the current user:

```sh
gator register alice
```

Add an RSS feed. This also makes the current user follow it:

```sh
gator addfeed "Example Blog" "https://example.com/rss.xml"
```

Start the feed aggregator, choosing how often it checks for updates. This command continues running; leave it open while posts are collected:

```sh
gator agg 1m
```

In another terminal, inspect feeds and browse collected posts:

```sh
gator feeds
gator following
gator browse 10
```

The duration uses Go's duration syntax, such as `30s`, `1m`, or `1h`.

## Commands

| Command | Description |
| --- | --- |
| `register <username>` | Create a user and set them as the current user. |
| `login <username>` | Set the current user to an existing account. |
| `users` | List registered users. |
| `addfeed <name> <url>` | Add a feed and follow it as the current user. |
| `feeds` | List feeds. |
| `follow <url>` | Follow an existing feed as the current user. |
| `following` | List feeds followed by the current user. |
| `unfollow <url>` | Stop following a feed. |
| `agg <duration>` | Fetch feeds repeatedly at the given interval; keeps running until stopped. |
| `browse [limit]` | Show recent posts for the current user's feeds (defaults to 2). |
| `reset` | Delete all users and associated data from the database. This is destructive. |

Commands that operate on a user's follows or posts require a current user. Use `register` or `login` first.

## Development

Run the test suite or build all packages from the repository root:

```sh
go test ./...
go build ./...
```

