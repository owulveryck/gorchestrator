package toscalib

import (
	"github.com/owulveryck/gorchestrator/Godeps/_workspace/src/github.com/gonum/matrix/mat64"
	"github.com/owulveryck/gorchestrator/Godeps/_workspace/src/gopkg.in/yaml.v2"
)

// ServiceTemplateDefinition is the meta structure containing an entire tosca document as described in
//http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
type ServiceTemplateDefinition struct {
	DefinitionsVersion Version                         `yaml:"tosca_definitions_version" json:"tosca_definitions_version"` // A.9.3.1 tosca_definitions_version
	Description        string                          `yaml:"description,omitempty" json:"description,omitempty"`
	Imports            []string                        `yaml:"imports,omitempty" json:"imports,omitempty"`                       // Declares import statements external TOSCA Definitions documents. For example, these may be file location or URIs relative to the service template file within the same TOSCA CSAR file.
	Repositories       map[string]RepositoryDefinition `yaml:"repositories,omitempty" json:"repositories,omitempty"`             // Declares the list of external repositories which contain artifacts that are referenced in the service template along with their addresses and necessary credential information used to connect to them in order to retrieve the artifacts.
	DataTypes          map[string]DataType             `yaml:"data_types,omitempty" json:"data_types,omitempty"`                 // Declares a list of optional TOSCA Data Type definitions.
	NodeTypes          map[string]NodeType             `yaml:"node_types,omitempty" json:"node_types,omitempty"`                 // This section contains a set of node type definitions for use in service templates.
	RelationshipTypes  map[string]RelationshipType     `yaml:"relationship_types,omitempty" json:"relationship_types,omitempty"` // This section contains a set of relationship type definitions for use in service templates.
	CapabilityTypes    map[string]CapabilityType       `yaml:"capability_types,omitempty" json:"capability_types,omitempty"`     // This section contains an optional list of capability type definitions for use in service templates.
	ArtifactTypes      map[string]ArtifactType         `yaml:"artifact_types,omitempty" json:"artifact_types,omitempty"`         // This section contains an optional list of artifact type definitions for use in service templates
	DlsDefinitions     interface{}                     `yaml:"dsl_definitions,omitempty" json:"dsl_definitions,omitempty"`       // Declares optional DSL-specific definitions and conventions.  For example, in YAML, this allows defining reusable YAML macros (i.e., YAML alias anchors) for use throughout the TOSCA Service Template.
	InterfaceTypes     map[string]InterfaceType        `yaml:"interface_types,omitempty" json:"interface_types,omitempty"`       // This section contains an optional list of interface type definitions for use in service templates.
	TopologyTemplate   TopologyTemplateType            `yaml:"topology_template" json:"topology_template"`                       // Defines the topology template of an application or service, consisting of node templates that represent the application’s or service’s components, as well as relationship templates representing relations between the components.
	AdjacencyMatrix    mat64.Dense                     `yaml:"-"`                                                                //The AdjacencyMatrix
}

// Bytes returns the yaml representation of the current definition in bytes
func (toscaDefinition *ServiceTemplateDefinition) Bytes() []byte {
	d, _ := yaml.Marshal(toscaDefinition)
	return d
}

// String returns the yaml representation of the current definition as a string
func (toscaDefinition *ServiceTemplateDefinition) String() string {
	d, _ := yaml.Marshal(toscaDefinition)
	return string(d)
}
