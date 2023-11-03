package api

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Log *logrus.Logger

func InitLogger() {
	Log = logrus.New()
	Log.Out = os.Stdout
	Log.SetFormatter(&logrus.JSONFormatter{})
}
