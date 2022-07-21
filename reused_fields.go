package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("\n========================== Reuse Fields ==========================")

	// save logger entry
	loggerEntry := log.WithFields(log.Fields{
		"user_id":  "tmp_user",
		"passcode": 987,
	})

	loggerEntry.Info("I'll be logged with two reusable fields")
	loggerEntry.Warn("Me too")

	// extend the existing logger entry
	loggerEntry = loggerEntry.WithField("new_key", "new_value")

	loggerEntry.Error("I'll be logged with three reusable fields")
}
