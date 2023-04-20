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

var _ = Describe("File Configuration", func() {
	Context("When the specified path does not exist", func() {
		config, err := hermes.NewFileConfiguration("/does/not/exist.conf")

		It("returns an ErrNotExist error", func() {
			Expect(errors.Is(err, os.ErrNotExist)).To(BeTrue())
		})

		It("returns a nil configuration", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a valid file is provided", func() {
		config, err := hermes.NewFileConfiguration("./data/valid.conf")

		It("has a nil error", func() {
			Expect(err).To(BeNil())
		})

		It("does not return a nil Configuration", func() {
			Expect(config).To(Not(BeNil()))
		})

		It("has the Sender populated", func() {
			Expect(config.Sender).To(Equal("mysender"))
		})

		It("has the Receiver populated", func() {
			Expect(config.Receiver).To(Equal("myreceiver"))
		})

		It("has the Mail Host populated", func() {
			Expect(config.MailHost).To(Equal("mymailhost"))
		})

		It("has the Mail Port populated", func() {
			Expect(config.MailPort).To(Equal("mymailport"))
		})

		It("has the Password populated", func() {
			Expect(config.Passwd).To(Equal("secret"))
		})
	})

	Context("When a file is missing the Sender", func() {
		config, err := hermes.NewFileConfiguration("./data/missing_sender.conf")

		It("returns Sender not set error", func() {
			Expect(errors.Is(err, hermes.ErrSenderNotSet)).To(BeTrue())
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a file is missing the Receiver", func() {
		config, err := hermes.NewFileConfiguration("./data/missing_receiver.conf")

		It("returns Receiver not set error", func() {
			Expect(errors.Is(err, hermes.ErrReceiverNotSet)).To(BeTrue())
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a file is missing the Mail Host", func() {
		config, err := hermes.NewFileConfiguration("./data/missing_mailhost.conf")

		It("returns Mail Host not set error", func() {
			Expect(errors.Is(err, hermes.ErrMailHostNotSet)).To(BeTrue())
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a file is missing the Mail Port", func() {
		config, err := hermes.NewFileConfiguration("./data/missing_mailport.conf")

		It("returns Mail Port not set error", func() {
			Expect(errors.Is(err, hermes.ErrMailPortNotSet)).To(BeTrue())
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a file is missing the Password", func() {
		config, err := hermes.NewFileConfiguration("./data/missing_password.conf")

		It("returns Password not set error", func() {
			Expect(errors.Is(err, hermes.ErrPasswdNotSet)).To(BeTrue())
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})

	Context("When a file contains invalid JSON", func() {
		config, err := hermes.NewFileConfiguration("./data/invalid_json.conf")

		It("returns an error", func() {
			Expect(err).To(Not(BeNil()))
		})

		It("has a nil config", func() {
			Expect(config).To(BeNil())
		})
	})
})

var _ = Describe("Environment Variable Configuration", func() {

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
			config, err = hermes.NewEnvConfiguration()
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
			config, err = hermes.NewEnvConfiguration()
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
			config, err = hermes.NewEnvConfiguration()
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
			config, err = hermes.NewEnvConfiguration()
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
			config, err = hermes.NewEnvConfiguration()
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
			config, err = hermes.NewEnvConfiguration()
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
