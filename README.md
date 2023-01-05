# nsddyn

Hermes provides a mechanism to notify someone about cronjob failures.

## Installing and Getting Started

Hermes is currently a source-only distribution. Both a Unix-like system and Google Go is required to perform compilation.
Check the go.mod file for the minimum required Go version.

### Compilation

1. Clone the repo.
2. `cd hermes/cmd/hermes`
3. `go build`

### Installation and Configuration

The following environment variables need to be defined and placed into a file such as `/etc/hermes.env`

1. `HERMES_SENDER="notifier@example.com"`
2. `HERMES_RECEIVER="notifiee@example.com"`
3. `HERMES_HOST="mail.example.com"`
4. `HERMES_PORT="587"`
5. `HERMES_PASSWD="notifiers_email_password"`

Be sure to protect `hermes.env` via `chown root:root hermes.env && chmod 0400 hermes.env`.

Copy the `hermes` executable into `/usr/local/bin/` or somewhere reasonable.

Edit `/etc/cronjob` so that jobs which require notification use `hermes` as an OR branch.
For example: \
`0  4    15 * *  root    zpool scrub zp0 zp1` \
Becomes: \
`0  4    15 * *  root    (zpool scrub zp0 zp1) || (. /etc/hermes.env; hermes -m "Failed to initial zpool scrub!!!")`

### Executing Tests

Hermes requires Ginkgo for testing but tests can be executed with `go test` from the root of the repo.
Although, it is recommended to use `ginkgo` explicitly rather than `go test` for test execution.

## Licensing

Hermes is released under the terms of the GPLv3 license, a copy of the GPL is provided in the COPYING
file located at the root of this repo.
