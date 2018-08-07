package client

import (
	"runtime"
	"os/exec"
	"fmt"
	"github.com/sirupsen/logrus"
)

func NewNotification(sendNotification bool, log *logrus.Logger) func(string, string) {
	if sendNotification {
		log.Debug("Send desktop notification has been enabled.")

		return func(title, text string) {
			switch runtime.GOOS {
			case "darwin":
				exec.Command("osascript", "-e", fmt.Sprintf("display notification \"%s\" with title \"%s\"", text, title)).Run()
			case "linux":
				exec.Command("notify-send", title, text).Run()
			default:
				// [TODO] Sending notification must be implemented for windows or use a library of golang which do that!
			}
		}
	}

	return func(title, text string) {}
}
