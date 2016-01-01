package toscalib

// Input corresponds to  `yaml:"inputs,omitempty" json:"inputs,omitempty"`
type Input struct {
	Type             string      `yaml:"type" json:"type"`
	Description      string      `yaml:"description,omitempty" json:"description,omitempty"` // Not required
	Constraints      Constraints `yaml:"constraints,omitempty" json:"constraints,omitempty"`
	ValidSourceTypes interface{} `yaml:"valid_source_types,omitempty" json:"valid_source_types,omitempty"`
	Occurrences      interface{} `yaml:"occurrences,omitempty" json:"occurrences,omitempty"`
}
