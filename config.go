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
	"encoding/json"
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

const (
	configPrefix = "/opt/catloaf/etc"
	configFile   = "hermes.conf"
)

type Configuration struct {
	Sender   string
	Receiver string
	MailHost string
	MailPort string
	Passwd   string
}

type FileOptions struct {
	Sender   string
	Receiver string
	Host     string
	Port     string
	Password string
}

// NewConfiguration is the default entry point for getting the hermes Configuration.
// Unless you have a good reason to do otherwise, NewConfiguration should be used over other methods.
// NewConfiguration uses a precedence when evaluating configuration mechanisms, the first to succeed wins.
// If no mechanism is able to succeed, Configuration is undefined and err != nil.
// Mechanism precedence is as follows:
// 1. Configuration file in the default location. [NewDefaultFileConfiguration()]
// 2. Environment variables. [NewEnvConfiguration()]
func NewConfiguration() (*Configuration, error) {
	_, err := os.Stat(configPrefix + "/" + configFile)
	if os.IsNotExist(err) {
		return NewEnvConfiguration()
	}
	return NewDefaultFileConfiguration()
}

// NewDefaultFileConfiguration returns a hermes Configuration with options based on a configuration file in the default location.
// If there is a problem reading the file, the returned Configuration is undefined and err != nil.
// The default location of the config file is: /opt/catloaf/etc/hermes.conf
func NewDefaultFileConfiguration() (*Configuration, error) {
	configFilePath := configPrefix + "/" + configFile
	return NewFileConfiguration(configFilePath)
}

// NewFileConfiguration returns a hermes Configuration with options based on the specified configuration file.
// If there is a problem reading the file, the returned Configuration is undefined and err != nil.
func NewFileConfiguration(configFilePath string) (*Configuration, error) {
	_, err := os.Stat(configFilePath)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}

	fo := FileOptions{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&fo)
	if err != nil {
		return nil, err
	}

	config := new(Configuration)

	if fo.Sender == "" {
		return nil, ErrSenderNotSet
	}
	config.Sender = fo.Sender

	if fo.Receiver == "" {
		return nil, ErrReceiverNotSet
	}
	config.Receiver = fo.Receiver

	if fo.Host == "" {
		return nil, ErrMailHostNotSet
	}
	config.MailHost = fo.Host

	if fo.Port == "" {
		return nil, ErrMailPortNotSet
	}
	config.MailPort = fo.Port

	if fo.Password == "" {
		return nil, ErrPasswdNotSet
	}
	config.Passwd = fo.Password

	return config, nil
}

// NewEnvConfiguration returns a hermes Configuration with options based on environment variables.
// If any one of the environment variables is missing, the returned Configuration is undefined and err != nil.
// The following environment variables are used:
// HERMES_SENDER
// HERMES_RECEIVER
// HERMES_HOST
// HERMES_PORT
// HERMES_PASSWD
func NewEnvConfiguration() (*Configuration, error) {
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
