/*
   Copyright (C) 2023, 2024 Alexander Necheff

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

package hermes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"necheff.net/hermes"

	"errors"
)

type MockConnection struct {
	hermes.Configuration
}

func NewMockConnection(config *hermes.Configuration) *MockConnection {
	c := new(MockConnection)

	return c
}

func (mc *MockConnection) To() string {
	return mc.Receiver
}

func (mc *MockConnection) From() string {
	return mc.Sender
}

func (mc *MockConnection) SendMessage(msg string) error {
	_ = msg
	return errors.New("a junk error")
}

var _ = Describe("Hermes", func() {
	Context("When a send message has an error", func() {
		config := &hermes.Configuration{
			Sender:   "Sender",
			Receiver: "Receiver",
			MailHost: "Host",
			MailPort: "Port",
			Passwd:   "Passwd",
		}
		conn := NewMockConnection(config)
		err := hermes.Hermes(conn, "fake message")

		It("returns an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})
})
