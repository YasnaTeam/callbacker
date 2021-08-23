package logger

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

//var log *Logger
//var log = &Logger{
//	&logrus.Logger{
//		//Formatter: &easy.Formatter{
//		//	TimestampFormat: "2006-01-02 15:04:05.99999",
//		//	LogFormat:       "%lvl%[%time%] %msg%\n",
//		//},
//	},
//}
var log = logrus.New()

func New(debug bool) *logrus.Logger {
	if debug {
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	} else {
		log.SetLevel(logrus.FatalLevel)
	}

	return log
}

func WithRequest(logger *logrus.Logger, r *http.Request) *logrus.Entry {
	return logger.WithFields(logrus.Fields{
		"method":      r.Method,
		"protocol":    r.Proto,
		"remote_addr": r.RemoteAddr,
	})
}
