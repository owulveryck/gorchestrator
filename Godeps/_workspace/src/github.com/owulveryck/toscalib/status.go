package toscalib

// Status is used in the PropertyDefinition
type Status string

// Valid values for Status as described in Appendix 5.7.3
const (
	Supported    Status = "supported"
	Unsupported  Status = "unsupported"
	Experimental Status = "experimental"
	Deprecated   Status = "deprecated"
)
