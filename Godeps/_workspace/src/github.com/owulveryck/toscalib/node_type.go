package toscalib

import (
	"fmt"
)

// NodeType as described is Appendix 6.8.
// A Node Type is a reusable entity that defines the type of one or more Node Templates. As such, a Node Type defines the structure of observable properties via a Properties Definition, the Requirements and Capabilities of the node as well as its supported interfaces.
type NodeType struct {
	DerivedFrom  string                          `yaml:"derived_from,omitempty" json:"derived_from"`           // An optional parent Node Type name this new Node Type derives from
	Description  string                          `yaml:"description,omitempty" json:"description"`             // An optional description for the Node Type
	Properties   map[string]PropertyDefinition   `yaml:"properties,omitempty" json:"properties,omitempty"`     // An optional list of property definitions for the Node Type.
	Attributes   map[string]AttributeDefinition  `yaml:"attributes,omitempty" json:"attributes,omitempty"`     // An optional list of attribute definitions for the Node Type.
	Requirements []map[string]interface{}        `yaml:"requirements,omitempty" json:"requirements,omitempty"` // An optional sequenced list of requirement definitions for the Node Type
	Capabilities map[string]CapabilityDefinition `yaml:"capabilities,omitempty" json:"capabilities,omitempty"` // An optional list of capability definitions for the Node Type
	Interfaces   map[string]InterfaceDefinition  `yaml:"interfaces,omitempty" json:"interfaces,omitempty"`     // An optional list of interface definitions supported by the Node Type
	Artifacts    []ArtifactDefinition            `yaml:"artifacts,omitempty" json:"artifacts,omitempty"`       // An optional list of named artifact definitions for the Node Type
	Copy         string                          `yaml:"copy,omitempty" json:"copy,omitempty"`                 // The optional (symbolic) name of another node template to copy into (all keynames and values) and use as a basis for this node template.
}

func (n *NodeType) getInterface() (string, InterfaceDefinition, error) {
	for name, value := range n.Interfaces {
		return name, value, nil
	}
	return "", InterfaceDefinition{}, fmt.Errorf("No Interface found")
}
