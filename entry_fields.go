package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

//func init() {
//	log.SetFormatter(&log.JSONFormatter{})
//}

func main() {
	showBuiltInFields()
	addFields()
	addError()
}

func showBuiltInFields() {
	fmt.Println("\n========================== Show Built-in Fields ==========================")

	log.Info("This only has built-in fields")
}

func addFields() {
	fmt.Println("\n========================== Add Fields ==========================")

	log.WithField("any_key", "any_value").
		Warn("One extra field added")

	log.WithFields(log.Fields{
		"location": "movie theater",
		"food":     "popcorn",
		"people":   30,
	}).Error("Multiple fields added")
}

func addError() {
	fmt.Println("\n========================== Add Error ==========================")

	err := errors.New("an error")
	log.WithError(err).Error("Failed to xxx")
}
