package toscalib

import (
	"fmt"
	"regexp"
)

// NodeTemplate as described in Appendix 7.3
// A Node Template specifies the occurrence of a manageable software component as part of an application’s topology model which is defined in a TOSCA Service Template.  A Node template is an instance of a specified Node Type and can provide customized properties, constraints or operations which override the defaults provided by its Node Type and its implementations.
type NodeTemplate struct {
	Name         string
	Type         string                             `yaml:"type" json:"type"`                                              // The required name of the Node Type the Node Template is based upon.
	Decription   string                             `yaml:"description,omitempty" json:"description,omitempty"`            // An optional description for the Node Template.
	Directives   []string                           `yaml:"directives,omitempty" json:"-" json:"directives,omitempty"`     // An optional list of directive values to provide processing instructions to orchestrators and tooling.
	Properties   map[string]PropertyAssignment      `yaml:"properties,omitempty" json:"-" json:"properties,omitempty"`     // An optional list of property value assignments for the Node Template.
	Attributes   map[string]interface{}             `yaml:"attributes,omitempty" json:"-" json:"attributes,omitempty"`     // An optional list of attribute value assignments for the Node Template.
	Requirements []map[string]RequirementAssignment `yaml:"requirements,omitempty" json:"-" json:"requirements,omitempty"` // An optional sequenced list of requirement assignments for the Node Template.
	Capabilities map[string]interface{}             `yaml:"capabilities,omitempty" json:"-" json:"capabilities,omitempty"` // An optional list of capability assignments for the Node Template.
	Interfaces   map[string]InterfaceType           `yaml:"interfaces,omitempty" json:"-" json:"interfaces,omitempty"`     // An optional list of named interface definitions for the Node Template.
	Artifcats    map[string]ArtifactDefinition      `yaml:"artifcats,omitempty" json:"-" json:"artifcats,omitempty"`       // An optional list of named artifact definitions for the Node Template.
	NodeFilter   map[string]NodeFilter              `yaml:"node_filter,omitempty" json:"-" json:"node_filter,omitempty"`   // The optional filter definition that TOSCA orchestrators would use to select the correct target node.  This keyname is only valid if the directive has the value of “selectable” set.
	Refs         struct {
		Type       *NodeType        `yaml:"-",json:"-"`
		Interfaces []*InterfaceType `yaml:"-",json:"-"`
	} `yaml:"-",json:"-"`
}

// setRefs fills in the references of the node
func (n *NodeTemplate) setRefs(s *ServiceTemplateDefinition) {
	for name, _ := range n.Interfaces {
		re := regexp.MustCompile(fmt.Sprintf("^%v$", name))
		for na, v := range s.InterfaceTypes {
			if re.MatchString(na) {
				n.Refs.Interfaces = append(n.Refs.Interfaces, &v)
			}
		}
	}
	for na, v := range s.NodeTypes {
		if na == n.Type {
			n.Refs.Type = &v
		}
	}
}

func (n *NodeTemplate) getInterface() (string, InterfaceType, error) {
	for name, value := range n.Interfaces {
		return name, value, nil
	}
	return "", InterfaceType{}, fmt.Errorf("No Interface found")
}

// fillInterface Completes the interface of the node with any values found in its type
// All the Operations will be filled
func (n *NodeTemplate) fillInterface(s ServiceTemplateDefinition) {
	name, intf, err := n.getInterface()
	if err != nil {
		return
	}
	nt := s.NodeTypes[n.Type]
	_, intf2, _ := nt.getInterface()
	re := regexp.MustCompile(fmt.Sprintf("%v$", name))
	operations := make(map[string]OperationDefinition, 0)
	for ifacename, iface := range s.InterfaceTypes {
		if re.MatchString(ifacename) {
			for op, _ := range iface.Operations {
				v, ok := intf.Operations[op]
				_, ok2 := intf2[op]
				switch {
				case !ok && ok2:
					operations[op] = OperationDefinition{nil, intf2[op].Description, intf2[op].Implementation}
				case ok:
					operations[op] = v
				default:
				}
				tmp := InterfaceType{n.Interfaces[name].Description, n.Interfaces[name].Version, operations, n.Interfaces[name].Inputs}
				n.Interfaces[name] = tmp
			}
		}
	}
}

func (n *NodeTemplate) setName(name string) {
	n.Name = name
}
