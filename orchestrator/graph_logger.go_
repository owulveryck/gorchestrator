package orchestrator

import (
	"github.com/owulveryck/gorchestrator/logger"
)

func (g *Graph) getFields() logger.Fields {
	return logger.Fields{
		"Element": "Graph",
		"ID":      g.ID,
		"State":   g.State,
	}
}

// Log an element with all the fields of the node
func (g *Graph) Log(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Println(args...)
}

// Log an element with all the fields of the node
func (g *Graph) Logf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Printf(format, args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogDebug(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Debug(args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogDebugf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Debugf(format, args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogWarn(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Warn(args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogWarnf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Warnf(format, args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogError(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Error(args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogErrorf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Errorf(format, args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogInfo(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Info(args...)
}

// Log an element with all the fields of the node
func (g *Graph) LogInfof(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(g.getFields().Fields()).Infof(format, args...)
}
