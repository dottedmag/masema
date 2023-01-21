# Masema

Masema is a Mastodon to e-mail gateway.

* Sends messages from a home timeline
* Shows pictures inline

# Installation and configuration

* Build the tool (e.g. `go install github.com/dottedmag/masema@latest`
* Configure local mail server to accept mail via `/usr/sbin/sendmail`
* Copy [example config](docs/config.toml.example) to `$HOME/.config/masema/config.toml`
* Fill in `mail.from` and `mail.to` values
* [Register an app](https://docs.joinmastodon.org/client/token/#app)
* Fill in `app.server_url` and `app.token` values

# Usage

Fetch messages:
```
$ masema fetch
```

Send messages:
```
$ masema send
```
