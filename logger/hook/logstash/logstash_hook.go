package logstash

import (
	"github.com/Sirupsen/logrus"
	logstash "github.com/Sirupsen/logrus/formatters/logstash"
	"net"
)

type LogstashHook struct {
	conn        net.Conn
	application string
	level       logrus.Level
}

func NewLogstashHook(protocol, address, application string, level logrus.Level) (*LogstashHook, error) {
	conn, err := net.Dial(protocol, address)
	if err != nil {
		return nil, err
	}
	return &LogstashHook{conn: conn, application: application, level: level}, nil
}

func (h *LogstashHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func (h *LogstashHook) Fire(entry *logrus.Entry) error {
	if entry.Level >= h.level {
		formatter := logstash.LogstashFormatter{Type: h.application}
		dataBytes, err := formatter.Format(entry)
		if err != nil {
			return err
		}
		if _, err = h.conn.Write(dataBytes); err != nil {
			return err
		}
	}
	return nil
}
