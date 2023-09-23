# hermes

Hermes provides a mechanism to notify someone about cronjob failures.

## Installing and Getting Started

Hermes is currently a source-only distribution. Both a Unix-like system and Google Go is required to perform compilation.
Check the go.mod file for the minimum required Go version.

### Compilation

1. Clone the repo.
2. `cd hermes && make`
3. The `cmd/hermes/hermes` binary should now exist.

### Installation

The `hermes` executable can be simply `cp`'d into `/usr/local/bin/` or somewhere reasonable.

However, a `debian` target exists in the Makefile which will generate a Debian package provided your system is able to execute `dpkg-deb`.

### Configuration

Currently hermes supports both a file and an environment variable configuration mechanism.
The prefered way is to use the file mechanism.

#### File Configuration Mechanism

Currently, there is no way to override the default configuration file location, which is: `/opt/catloaf/etc/hermes.conf`.
This file is a JSON file and if it exists, hermes will prefer its settings. The following format is used:

```
{
"Sender":"mysender",
"Receiver":"myreceiver",
"Host":"mymailhost",
"Port":"mymailport",
"Password":"secret"
}
```

To protect the file, a hermes user needs to be created to own the file. Make the `hermes` binary sgid to allow any user with
shell access to send messages with hermes. Or, don't set the sgid bit and require membership to the hermes group to limit shell user
access to the messaging capabilities provided by hermes.

1. First, create the user and group: `adduser --disabled-login --no-create-home --uid $UID hermes`
2. Next, take ownership of the binary and config file: `chown root:hermes /usr/local/bin/hermes && chown root:hermes /opt/catloaf/etc/hermes.conf"`
3. Lastly, if desired, set the sgid bit: `chmod 2555 /usr/local/bin/hermes`. Or, require group membership instead: `chmod 0550 /usr/local/bin/hermes`.

Edit `/etc/crontab` so that jobs which require notification use `hermes` as an OR branch.
For example: \
`0  4    15 * *  root    zpool scrub zp0 zp1` \
Becomes: \
`0  4    15 * *  root    (zpool scrub zp0 zp1) || hermes -m "Failed to initalize zpool scrub!!!"`

#### Environment Variables Configuration Mechanism

The following environment variables need to be defined and placed into a file such as `/etc/hermes.env`

1. `HERMES_SENDER="notifier@example.com"`
2. `HERMES_RECEIVER="notifiee@example.com"`
3. `HERMES_HOST="mail.example.com"`
4. `HERMES_PORT="587"`
5. `HERMES_PASSWD="notifiers_email_password"`

Be sure to protect `hermes.env` via `chown root:root hermes.env && chmod 0400 hermes.env`.

Edit `/etc/cronjob` so that jobs which require notification use `hermes` as an OR branch.
For example: \
`0  4    15 * *  root    zpool scrub zp0 zp1` \
Becomes: \
`0  4    15 * *  root    (zpool scrub zp0 zp1) || (. /etc/hermes.env; hermes -m "Failed to initialize zpool scrub!!!")`

### Executing Tests

Hermes requires Ginkgo for testing. The `test` target in the Makefile is the recommended way to execute the test suite.
This will also display coverage metrics automatically.

## Licensing

Hermes is released under the terms of the GPLv3 license, a copy of the GPL is provided in the COPYING
file located at the root of this repo.
