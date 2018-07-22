package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"github.com/t-tomalak/logrus-easy-formatter"
	"net/http"
)

var log = &logrus.Logger {
	Formatter: &easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05.99999",
		LogFormat:       "%lvl%[%time%] %msg%\n",
	},
}

func New(debug bool) *logrus.Logger {
	if debug {
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	return log
}

func WithRequest(r *http.Request) *logrus.Entry {
	return log.WithFields(requestFields(r))
}

func requestFields(r *http.Request) logrus.Fields {
	return logrus.Fields{
		"method": r.Method,
		"host": r.Host,
	}
}