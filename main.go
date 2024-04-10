package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/uwezukwechibuzor/bitcoin-node-handshake/cmd"
)

func main() {
	if os.Getenv("DEBUG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cmd.Execute()
}
