package client

import (
	"github.com/sirupsen/logrus"
)


func Initialize(serverAddress string, log *logrus.Logger) {
	log.Info("Client started...")
}
