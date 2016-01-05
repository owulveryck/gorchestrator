package toscalib

// CapabilityDefinition TODO: Appendix 6.1
type CapabilityDefinition struct {
	Type             string                `yaml:"type" json:"type"`                                    //  The required name of the Capability Type the capability definition is based upon.
	Description      string                `yaml:"description,omitempty" jsson:"description,omitempty"` // The optional description of the Capability definition.
	Properties       []PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`    //  An optional list of property definitions for the Capability definition.
	Attributes       []AttributeDefinition `yaml:"attributes" json:"attributes"`                        // An optional list of attribute definitions for the Capability definition.
	ValidSourceTypes []string              `yaml:"valid_source_types" json:"valid_source_types"`        // A`n optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type.
	Occurences       []string              `yaml:"occurences" json:"occurences"`
}

// UnmarshalYAML is used to match both Simple Notation Example and Full Notation Example
func (c *CapabilityDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try the Short notation
	var cas string
	err := unmarshal(&cas)
	if err == nil {
		c.Type = cas
		return nil
	}
	// If error, try the full struct
	type cap struct {
		Type             string                `yaml:"type" json:"type"`                                    //  The required name of the Capability Type the capability definition is based upon.
		Description      string                `yaml:"description,omitempty" jsson:"description,omitempty"` // The optional description of the Capability definition.
		Properties       []PropertyDefinition  `yaml:"properties,omitempty" json:"properties,omitempty"`    //  An optional list of property definitions for the Capability definition.
		Attributes       []AttributeDefinition `yaml:"attributes" json:"attributes"`                        // An optional list of attribute definitions for the Capability definition.
		ValidSourceTypes []string              `yaml:"valid_source_types" json:"valid_source_types"`        // A`n optional list of one or more valid names of Node Types that are supported as valid sources of any relationship established to the declared Capability Type.
		Occurences       []string              `yaml:"occurences" json:"occurences"`
	}
	var ca cap
	err = unmarshal(&ca)
	if err != nil {
		return err
	}
	c.Type = ca.Type
	c.Description = ca.Description
	c.Properties = ca.Properties
	c.Attributes = ca.Attributes
	c.Occurences = ca.Occurences
	c.ValidSourceTypes = ca.ValidSourceTypes

	return nil
}

// CapabilityType as described in appendix 6.6
//A Capability Type is a reusable entity that describes a kind of capability that a Node Type can declare to expose.  Requirements (implicit or explicit) that are declared as part of one node can be matched to (i.e., fulfilled by) the Capabilities declared by another node.
// TODO
type CapabilityType interface{}
