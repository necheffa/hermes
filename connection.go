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
	"net/smtp"
)

const (
	MailNewLine = "\r\n"
)

type Connection interface {
	SendMessage(msg string) error
	To() string
	From() string
}

type SmtpConnection struct {
	auth smtp.Auth
	to   string
	from string
	host string
	port string
}

func NewSmtpConnection(config *Configuration) (*SmtpConnection, error) {
	c := new(SmtpConnection)
	c.to = config.Receiver
	c.from = config.Sender
	c.host = config.MailHost
	c.port = config.MailPort

	c.auth = smtp.PlainAuth("", config.Sender, config.Passwd, config.MailHost)

	return c, nil
}

func (conn *SmtpConnection) To() string {
	return conn.to
}

func (conn *SmtpConnection) From() string {
	return conn.from
}

func (conn *SmtpConnection) MailHost() string {
	return conn.host + ":" + conn.port
}

func (conn *SmtpConnection) SendMessage(msg string) error {
	mail := []byte("To: " + conn.To() + MailNewLine +
		"From: " + conn.From() + MailNewLine +
		"Subject: ALERT from hermes" + MailNewLine +
		MailNewLine +
		msg + MailNewLine)

	sendToList := []string{conn.To()}

	return smtp.SendMail(conn.MailHost(), conn.auth, conn.From(), sendToList, mail)
}
