/*
   Copyright (C) 2023 Alexander Necheff

   This file is part of hermes.

   hermes is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, at version 3 of the License.

   hermes is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with hermes.  If not, see <https://www.gnu.org/licenses/>.
*/

package hermes

import (
	"errors"
	"os"
)

var (
	ErrSenderNotSet   = errSenderNotSet()
	ErrReceiverNotSet = errReceiverNotSet()
	ErrMailHostNotSet = errMailHostNotSet()
	ErrMailPortNotSet = errMailPortNotSet()
	ErrPasswdNotSet   = errPasswdNotSet()
)

const (
	Sender   = "HERMES_SENDER"
	Receiver = "HERMES_RECEIVER"
	Host     = "HERMES_HOST"
	Port     = "HERMES_PORT"
	Passwd   = "HERMES_PASSWD"
)

type Configuration struct {
	Sender   string
	Receiver string
	MailHost string
	MailPort string
	Passwd   string
}

func NewConfiguration() (*Configuration, error) {
	config := new(Configuration)

	sender, ok := os.LookupEnv(Sender)
	if !ok {
		return nil, ErrSenderNotSet
	}
	config.Sender = sender

	receiver, ok := os.LookupEnv(Receiver)
	if !ok {
		return nil, ErrReceiverNotSet
	}
	config.Receiver = receiver

	mailHost, ok := os.LookupEnv(Host)
	if !ok {
		return nil, ErrMailHostNotSet
	}
	config.MailHost = mailHost

	mailPort, ok := os.LookupEnv(Port)
	if !ok {
		return nil, ErrMailPortNotSet
	}
	config.MailPort = mailPort

	passwd, ok := os.LookupEnv(Passwd)
	if !ok {
		return nil, ErrPasswdNotSet
	}
	config.Passwd = passwd

	return config, nil
}

func errSenderNotSet() error {
	return errors.New("HERMES: error: $HERMES_SENDER not set")
}

func errReceiverNotSet() error {
	return errors.New("HERMES: error: $HERMES_RECEIVER not set")
}

func errMailHostNotSet() error {
	return errors.New("HERMES: error: $HERMES_HOST not set")
}

func errMailPortNotSet() error {
	return errors.New("HERMES: error: $HERMES_PORT not set")
}

func errPasswdNotSet() error {
	return errors.New("HERMES: error: $HERMES_PASSWD not set")
}
