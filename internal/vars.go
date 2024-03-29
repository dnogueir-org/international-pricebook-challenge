package internal

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var Logger = log.WithFields(log.Fields{
	"User":        "Daniel",
	"Application": "international-pricebook-challenge",
	"Environment": os.Getenv("ENV"),
})

func init() {
	log.SetFormatter(&log.JSONFormatter{})

	log.SetOutput(os.Stdout)

	log.SetLevel(log.WarnLevel)
}
