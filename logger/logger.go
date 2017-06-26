package logger

import (
	"os"
	"github.com/sirupsen/logrus"
)

var Log = *logrus.New()

func init() {
	Log.Out = os.Stdout
	Log.Level = logrus.InfoLevel
}
