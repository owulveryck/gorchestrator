package orchestrator

import (
	"github.com/owulveryck/gorchestrator/logger"
)

func (n *Node) getFields() logger.Fields {
	return logger.Fields{
		"Element":  "node",
		"ID":       n.ID,
		"GraphID":  n.GraphID,
		"ExecID":   n.execID,
		"State":    n.State,
		"Target":   n.Target,
		"Engine":   n.Engine,
		"Artifact": n.Artifact,
		"Args":     n.Args,
		"Outputs":  n.Outputs,
	}
}

// Log an element with all the fields of the node
func (n *Node) Log(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Println(args)
}

// Log an element with all the fields of the node
func (n *Node) Logf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Printf(format, args)
}

// Log an element with all the fields of the node
func (n *Node) LogDebug(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Debug(args)
}

// Log an element with all the fields of the node
func (n *Node) LogDebugf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Debugf(format, args)
}

// Log an element with all the fields of the node
func (n *Node) LogWarn(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Warn(args)
}

// Log an element with all the fields of the node
func (n *Node) LogWarnf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Warnf(format, args)
}

// Log an element with all the fields of the node
func (n *Node) LogError(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Error(args)
}

// Log an element with all the fields of the node
func (n *Node) LogErrorf(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Errorf(format, args)
}

// Log an element with all the fields of the node
func (n *Node) LogInfo(args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Info(args)
}

// Log an element with all the fields of the node
func (n *Node) LogInfof(format string, args ...interface{}) {
	log := logger.GetLog()
	log.WithFields(n.getFields().Fields()).Infof(format, args)
}
