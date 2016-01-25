package logger

import (
	"github.com/Sirupsen/logrus"
)

// This is a wrapper to the logrus.Entry to add the WithObject func
type Logger logrus.Entry

// The ObjectLogger interface
type ObjectLogger interface {
	GetFields() logrus.Fields
}

func (l *Logger) WithObject(o ObjectLogger) {
}
