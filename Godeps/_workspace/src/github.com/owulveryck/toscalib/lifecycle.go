package toscalib

// ToscaInterfacesNodeLifecycleStandarder is a go interface for the standard normative lifecycle
type ToscaInterfacesNodeLifecycleStandarder interface {
	Create() error    // description: Standard lifecycle create operation.
	Configure() error // description: Standard lifecycle configure operation.
	Start() error     // description: Standard lifecycle start operation.
	Stop() error      // description: Standard lifecycle stop operation.
	Delete() error    //description: Standard lifecycle delete operation.
}
