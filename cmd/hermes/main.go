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

package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"necheff.net/hermes"
)

//go:embed VERSION
var VERSION string

func main() {

	hermesFlags := flag.NewFlagSet("hermes", flag.ExitOnError)

	var printVersion bool
	var printHelp bool
	var message string

	hermesFlags.BoolVar(&printVersion, "version", false, "Display the version string and exit.")
	hermesFlags.BoolVar(&printVersion, "v", false, "Display the version string and exit.")
	hermesFlags.BoolVar(&printHelp, "help", false, "Display this usage message and exit.")
	hermesFlags.BoolVar(&printHelp, "h", false, "Display this usage message and exit.")
	hermesFlags.StringVar(&message, "message", "", "Specify the desired message body.")
	hermesFlags.StringVar(&message, "m", "", "Specify the desired message body.")

	err := hermesFlags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal("hermes: error: " + err.Error())
	}

	if printVersion {
		msg := "hermes v" + VERSION +
			"Built using library v" + hermes.Version +
			"Copyright (C) 2023\n" +
			"Alexander Necheff\n" +
			"hermes is licensed under the terms of the GPLv3."
		fmt.Println(msg)
		return
	}

	if printHelp {
		hermesFlags.PrintDefaults()
		return
	}

	config, err := hermes.NewConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := hermes.NewSmtpConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	err = hermes.Hermes(conn, message)
	if err != nil {
		log.Fatal(err)
	}
}
