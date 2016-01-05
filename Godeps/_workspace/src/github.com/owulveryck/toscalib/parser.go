package toscalib

import (
	"github.com/owulveryck/gorchestrator/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// GetNodeTemplate returns a pointer to a node template given its name
// its returns nil if not found
func (toscaStructure *ServiceTemplateDefinition) GetNodeTemplate(nodeName string) *NodeTemplate {
	for name, nodeTemplate := range toscaStructure.TopologyTemplate.NodeTemplates {
		if name == nodeName {
			return &nodeTemplate
		}
	}
	return nil
}

// merge copies all the elements of t into s and returns the result
func merge(s, t ServiceTemplateDefinition) ServiceTemplateDefinition {
	// Repositories
	rep := make(map[string]RepositoryDefinition, len(s.Repositories)+len(t.Repositories))
	for key, val := range t.Repositories {
		rep[key] = val
	}
	for key, val := range s.Repositories {
		rep[key] = val
	}
	s.Repositories = rep
	// DataTypes
	dat := make(map[string]DataType, len(s.DataTypes)+len(t.DataTypes))
	for key, val := range t.DataTypes {
		dat[key] = val
	}
	for key, val := range s.DataTypes {
		dat[key] = val
	}
	s.DataTypes = dat
	// NodeTypes
	nt := make(map[string]NodeType, len(s.NodeTypes)+len(t.NodeTypes))
	for key, val := range t.NodeTypes {
		nt[key] = val
	}
	for key, val := range s.NodeTypes {
		nt[key] = val
	}
	s.NodeTypes = nt
	// ArtifactType
	arti := make(map[string]ArtifactType, len(s.ArtifactTypes)+len(t.ArtifactTypes))
	for key, val := range t.ArtifactTypes {
		arti[key] = val
	}
	for key, val := range s.ArtifactTypes {
		arti[key] = val
	}
	s.ArtifactTypes = arti
	// RelationshipType
	rel := make(map[string]RelationshipType, len(s.RelationshipTypes)+len(t.RelationshipTypes))
	for key, val := range t.RelationshipTypes {
		rel[key] = val
	}
	for key, val := range s.RelationshipTypes {
		rel[key] = val
	}
	s.RelationshipTypes = rel
	// CapabilityType
	capa := make(map[string]CapabilityType, len(s.CapabilityTypes)+len(t.CapabilityTypes))
	for key, val := range t.CapabilityTypes {
		capa[key] = val
	}
	for key, val := range s.CapabilityTypes {
		capa[key] = val
	}
	s.CapabilityTypes = capa
	// InterfaceType
	intf := make(map[string]InterfaceType, len(s.InterfaceTypes)+len(t.InterfaceTypes))
	for key, val := range t.InterfaceTypes {
		intf[key] = val
	}
	for key, val := range s.InterfaceTypes {
		intf[key] = val
	}
	s.InterfaceTypes = intf
	return s
}

// Parse a TOSCA document and fill in the structure
func (t *ServiceTemplateDefinition) Parse(r io.Reader) error {
	var std ServiceTemplateDefinition
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	// Unmarshal the data in an interface
	err = yaml.Unmarshal(data, &std)
	if err != nil {
		return err
	}
	// Import de normative types by default
	for _, normType := range []string{"interface_types", "relationship_types", "node_types", "capability_types"} {
		data, err := Asset(normType)
		if err != nil {
			log.Panic("Normative type not found")
			return err
		}
		var tt ServiceTemplateDefinition
		err = yaml.Unmarshal(data, &tt)
		if err != nil {
			return err
		}
		std = merge(std, tt)
	}
	for _, im := range std.Imports {
		u, err := url.Parse(im)
		if err != nil {
			log.Panic(err)
		}
		var r []byte
		switch u.Scheme {
		case "http":
			res, err := http.Get(u.String())
			if err != nil {
				log.Fatal(err)

			}
			r, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)

			}
		case "https":
			res, err := http.Get(u.String())
			if err != nil {
				log.Fatal(err)

			}
			r, err = ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				log.Fatal(err)

			}
		default:
			r, err = ioutil.ReadFile(im)
			if err != nil {
				log.Fatal(err)
			}
		}
		var tt ServiceTemplateDefinition

		err = yaml.Unmarshal(r, &tt)
		if err != nil {
			log.Fatal(err)
		}
		std = merge(std, tt)
	}
	// Free the imports
	std.Imports = []string{}
	*t = std
	for name, node := range t.TopologyTemplate.NodeTemplates {
		node.fillInterface(*t)
		node.setRefs(t)
		node.setName(name)
	}

	return nil

}
