package toscalib

import (
	"fmt"
)

// PropertyDefinition as described in Appendix 5.7:
// A property definition defines a named, typed value and related data
// that can be associated with an entity defined in this specification
// (e.g., Node Types, Relation ship Types, Capability Types, etc.).
// Properties are used by template authors to provide input values to
// TOSCA entities which indicate their “desired state” when they are instantiated.
// The value of a property can be retrieved using the
// get_property function within TOSCA Service Templates
type PropertyDefinition struct {
	Type        string      `yaml:"type" json:"type"`                                   // The required data type for the property
	Description string      `yaml:"description,omitempty" json:"description,omitempty"` // The optional description for the property.
	Required    bool        `yaml:"required,omitempty" json:"required,omitempty"`       // An optional key that declares a property as required ( true) or not ( false) Default: true
	Default     string      `yaml:"default,omitempty" json:"default,omitempty"`
	Status      Status      `yaml:"status,omitempty" json:"status,omitempty"`
	Constraints Constraints `yaml:"constraints,omitempty,flow" json:"constraints,omitempty"`
	EntrySchema string      `yaml:"entry_schema,omitempty" json:"entry_schema,omitempty"`
}

// A Property assignment is always a map, but the key may be value
type PropertyAssignment map[string][]string

func (p *PropertyAssignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	*p = make(map[string][]string, 1)
	if err := unmarshal(&s); err == nil {
		(*p)["value"] = append([]string{}, s)
		return nil
	}
	var m map[string]string
	if err := unmarshal(&m); err != nil {
		for k, v := range m {
			(*p)[k] = append([]string{}, v)
		}
		return nil
	}
	var mm map[string][]string
	if err := unmarshal(&mm); err != nil {
		for k, v := range mm {
			(*p)[k] = v
		}
		return nil
	}
	var res interface{}
	unmarshal(&res)
	return fmt.Errorf("Cannot parse Property %v", res)
}
