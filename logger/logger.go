package logger

import (
	"github.com/Sirupsen/logrus"
	"github.com/Sirupsen/logrus/formatters/logstash"
	logrus_syslog "github.com/Sirupsen/logrus/hooks/syslog"
	"github.com/owulveryck/gorchestrator/config"
	"github.com/ripcurld00d/logrus-logstash-hook"
	"io"
	"log/syslog"
	"os"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.JSONFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

func init() {
	conf := config.GetConfig()
	if conf == nil {
		conf = &config.Config{}
	}
	log.Formatter = &logstash.LogstashFormatter{Type: "application_name"}
	// Output to stderr instead of stdout, could also be a file.
	var output io.Writer
	switch conf.Log.Output.Path {
	case "stderr":
		output = os.Stderr
	case "stdout":
		output = os.Stdout
	default:
		logrus.Warn("No log output defined, using stderr")
		output = os.Stderr
	}
	log.Out = output

	// Only log the warning severity or above.
	level, err := logrus.ParseLevel(conf.Log.Output.Level)
	if err != nil {
		logrus.Warn("Cannot parse conf level, default to INFO")
	} else {
		log.Level = level
	}
	for _, hook := range conf.Log.Hook {
		switch hook.Type {
		case "syslog":
			priority := syslog.LOG_INFO
			if err != nil {
				logrus.Errorf("%v is not a valid priority, defaulting to Info", hook.Level)
			}
			h, err := logrus_syslog.NewSyslogHook(hook.Protocol, hook.URL, priority, "")
			if err != nil {
				logrus.Errorf("Unable to connect to syslog daemon %v:%v", hook.Protocol, hook.URL)
			} else {
				log.Hooks.Add(h)
			}
		case "logstash":
			h, err := logrus_logstash.NewLogstashHook(hook.Protocol, hook.URL)
			if err != nil {
				logrus.Errorf("Unable to connect to logstash at %v:%v", hook.Protocol, hook.URL)

			} else {
				log.Hooks.Add(h)
			}
		}
	}
}

func GetLog() logrus.Logger {
	return *log
}
