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

package hermes_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"necheff.net/hermes"

	"errors"
	"os"
)

var _ = Describe("Configuration", func() {

	BeforeEach(func() {
		err := os.Setenv(hermes.Sender, "Sender")
		Expect(err).NotTo(HaveOccurred())

		err = os.Setenv(hermes.Receiver, "Receiver")
		Expect(err).NotTo(HaveOccurred())

		err = os.Setenv(hermes.Host, "Host")
		Expect(err).NotTo(HaveOccurred())

		err = os.Setenv(hermes.Port, "Port")
		Expect(err).NotTo(HaveOccurred())

		err = os.Setenv(hermes.Passwd, "Passwd")
		Expect(err).NotTo(HaveOccurred())
	})

	Context("When a Sender is not configured", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			os.Unsetenv(hermes.Sender)
			config, err = hermes.NewConfiguration()
		})

		It("returns sender not set error", func() {
			Expect(errors.Is(err, hermes.ErrSenderNotSet)).To(BeTrue())
		})
		It("has a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a Receiver is not configured", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			os.Unsetenv(hermes.Receiver)
			config, err = hermes.NewConfiguration()
		})

		It("returns Receiver not set error", func() {
			Expect(errors.Is(err, hermes.ErrReceiverNotSet)).To(BeTrue())
		})
		It("has a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a Host is not configured", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			os.Unsetenv(hermes.Host)
			config, err = hermes.NewConfiguration()
		})

		It("returns Host not set error", func() {
			Expect(errors.Is(err, hermes.ErrMailHostNotSet)).To(BeTrue())
		})
		It("has a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a Port is not configured", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			os.Unsetenv(hermes.Port)
			config, err = hermes.NewConfiguration()
		})

		It("returns Port not set error", func() {
			Expect(errors.Is(err, hermes.ErrMailPortNotSet)).To(BeTrue())
		})
		It("has a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a Passwd is not configured", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			os.Unsetenv(hermes.Passwd)
			config, err = hermes.NewConfiguration()
		})

		It("returns Passwd not set error", func() {
			Expect(errors.Is(err, hermes.ErrPasswdNotSet)).To(BeTrue())
		})
		It("has a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When all required configurations are specified", func() {
		var config *hermes.Configuration
		var err error

		JustBeforeEach(func() {
			config, err = hermes.NewConfiguration()
		})

		It("the returned error is nil", func() {
			Expect(err).To(BeNil())
		})

		It("the returned config is not nil", func() {
			Expect(config).To(Not(BeNil()))
		})

		It("the Sender is populated", func() {
			Expect(config.Sender).To(Equal("Sender"))
		})

		It("the Receiver is populated", func() {
			Expect(config.Receiver).To(Equal("Receiver"))
		})

		It("the Host is populated", func() {
			Expect(config.MailHost).To(Equal("Host"))
		})

		It("the Port is populated", func() {
			Expect(config.MailPort).To(Equal("Port"))
		})

		It("the Passwd is populated", func() {
			Expect(config.Passwd).To(Equal("Passwd"))
		})
	})
})
