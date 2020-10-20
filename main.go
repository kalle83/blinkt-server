package main

import (
	log "github.com/sirupsen/logrus"
	"vault.fritz.box/PI-ZERO-BLINKT-SERVER/cmd"
)

func main() {
	cmd.Execute()
	log.SetLevel(log.DebugLevel)

}