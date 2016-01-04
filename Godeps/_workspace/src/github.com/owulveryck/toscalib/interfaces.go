package toscalib

// InterfaceType as described in Appendix A 6.4
// An Interface Type is a reusable entity that describes a set of operations that can be used to interact with or manage a node or relationship in a TOSCA topology.
type InterfaceType struct {
	Description string                         `yaml:"description,omitempty"`
	Version     Version                        `yaml:"version,omitempty"`
	Operations  map[string]OperationDefinition `yaml:"operations,inline"`
	Inputs      map[string]PropertyDefinition  `yaml:"inputs,omitempty" json:"inputs"` // The optional list of input parameter definitions.
}

// InterfaceDefinition is related to a node type
//type InterfaceDefinitionTemplate map[string]OperationDefinition

// OperationDefinition defines a named function or procedure that can be bound to an implementation artifact (e.g., a script).
type OperationDefinition struct {
	Inputs         map[string]PropertyAssignment `yaml:"inputs,omitempty"`
	Description    string                        `yaml:"description,omitempty"`
	Implementation string                        `yaml:"implementation,omitempty"`
}

func (i *OperationDefinition) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs map[string]PropertyAssignment `yaml:"inputs,omitempty"`
		//Implementation      string                 `yaml:"implementation,omitempty"`
		Description    string `yaml:"description,omitempty"`
		Implementation string `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}

//type PropertyDefinition struct { }

// InterfaceDefinition TODO: Appendix 5.12

// InterfaceDefinition is related to a node type
type InterfaceDefinition map[string]InterfaceDef
type InterfaceDef struct {
	Inputs         map[string]Input `yaml:"inputs,omitempty"`
	Description    string           `yaml:"description,omitempty"`
	Implementation string           `yaml:"implementation,omitempty"`
}

func (i *InterfaceDef) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		i.Implementation = s
		return nil
	}
	var str struct {
		Inputs map[string]Input `yaml:"inputs,omitempty"`
		//Implementation      string                 `yaml:"implementation,omitempty"`
		Description    string `yaml:"description,omitempty"`
		Implementation string `yaml:"implementation,omitempty"`
	}
	if err := unmarshal(&str); err != nil {
		return err
	}
	i.Inputs = str.Inputs
	i.Implementation = str.Implementation
	i.Description = str.Description
	return nil
}
