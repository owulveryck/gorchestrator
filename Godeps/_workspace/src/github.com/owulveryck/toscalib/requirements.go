package toscalib

// RequirementDefinition as described in Appendix 6.2
type RequirementDefinition struct {
	Capability   string     `yaml:"capability" json:"capability"`                         // The required reserved keyname used that can be used to provide the name of a valid Capability Type that can fulfil the requirement
	node         string     `yaml:"node,omitempty" json:"node,omitempty"`                 // The optional reserved keyname used to provide the name of a valid Node Type that contains the capability definition that can be used to fulfil the requirement
	Relationship string     `yaml:"relationship,omitempty" json:"relationship,omitempty"` //The optional reserved keyname used to provide the name of a valid Relationship Type to construct when fulfilling the requirement.
	Occurrences  ToscaRange `yaml:"occurences,omitempty" json:"occurences,omitempty"`     // The optional minimum and maximum occurrences for the requirement.  Note: the keyword UNBOUNDED is also supported to represent any positive integer
}

// UnmarshalYAML is used to match both Simple Notation Example and Full Notation Example
func (r *RequirementAssignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// First try the Short notation
	var cas string
	err := unmarshal(&cas)
	if err == nil {
		r.Node = cas
		return nil
	}
	// If error, try the full struct
	type req struct {
		Capability   string      `yaml:"capability,omitempty"`
		Node         string      `yaml:"node,omitempty"`
		Nodefilter   NodeFilter  `yaml:"node_filter,omitempty"`
		Relationship interface{} `yaml:"relationship,omitempty"`
	}
	var ca req
	err = unmarshal(&ca)
	if err != nil {
		return err
	}
	r.Capability = ca.Capability
	r.Node = ca.Node
	r.Nodefilter = ca.Nodefilter
	r.Relationship = ca.Relationship
	return nil
}

// RequirementAssignment as described in Appendix 7.2
type RequirementAssignment struct {
	Capability string `yaml:"capability,omitempty" json:"capability,omitempty"` /* The optional reserved keyname used to provide the name of either a:
	- Capability definition within a target node template that can fulfill the requirement.
	- Capability Type that the provider will use to select a type-compatible target node template to fulfill the requirement at runtime.  */
	Node string `yaml:"node,omitempty" json:"node,omitempty"` /* The optional reserved keyname used to identify the target node of a relationship.  specifically, it is used to provide either a:
	-  Node Template name that can fulfil the target node requirement.
	- Node Type name that the provider will use to select a type-compatible node template to fulfil the requirement at runtime.  */
	//Relationship string `yaml:"relationship,omitempty" json:"relationship,omitempty"` /* The optional reserved keyname used to provide the name of either a:
	//- Relationship Template to use to relate the source node to the (capability in the) target node when fulfilling the requirement.
	//- Relationship Type that the provider will use to select a type-compatible relationship template to relate the source node to the target node at runtime. */
	Nodefilter NodeFilter `yaml:"node_filter,omitempty" json:"node_filter,omitempty"` // The optional filter definition that TOSCA orchestrators or providers would use to select a type-compatible target node that can fulfill the associated abstract requirement at runtime.o
	/* The following is the list of recognized keynames for a TOSCA requirement assignment’s relationship keyname which is used when Property assignments need to be provided to inputs of declared interfaces or their operations:*/
	Relationship interface{} `yaml:"relationship,omitempty" json:"relationship,omitempty"`
	// It looks like the Relationship type is not always present and from times to time (at least in the ELK example, we find the Interfaces directly)
}

/* The following is the list of recognized keynames for a TOSCA requirement assignment’s relationship keyname which is used when Property assignments need to be provided to inputs of declared interfaces or their operations:*/
type RequirementRelationship struct {
	Type       string                         `yaml:"type" json:"type"`                                 // The optional reserved keyname used to provide the name of the Relationship Type for the requirement assignment’s relationship keyname.
	Interfaces map[string]InterfaceDefinition `yaml:"interfaces,omitempty" json:"interfaces,omitempty"` // The optional reserved keyname used to reference declared (named) interface definitions of the corresponding Relationship Type in order to provide Property assignments for these interfaces or operations of these interfaces.
	Properties map[string]interface{}         `yaml:"properties" json:"properties"`                     // The optional list property definitions that comprise the schema for a complex Data Type in TOSCA.

}
