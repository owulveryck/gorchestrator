package logger

import (
	"github.com/Sirupsen/logrus"
	"testing"
)

func TestGetLog(t *testing.T) {
	a := GetLog()
	a.Debug("coucou")
	a.Info("coucou")
	a.Warning("coucou")
	a.Error("coucou")
	a.WithFields(logrus.Fields{"a": "b"}).Debug("blabla")
}
