package toscalib

// This implements the type defined in Appendix A 3 of the definition file
const (
	StateInitial     = iota // Node is not yet created. Node only exists as a template definition
	StateCreating    = iota // Node is transitioning from initial state to created state.
	StateCreated     = iota // Node software has been installed.
	StateConfiguring = iota // Node is transitioning from created state to configured state.
	StateConfigured  = iota // Node has been configured prior to being started
	StateStarting    = iota // Node is transitioning from configured state to started state.
	StateStarted     = iota // Node is started.
	StateStopping    = iota // Node is transitioning from its current state to a configured state.
	StateDeleting    = iota // Node is transitioning from its current state to one where it is deleted and its state is =iota // longer tracked by the instance model.
	StateError       = iota // Node is in an error state
)

// TODO: A 3.3 to describe
