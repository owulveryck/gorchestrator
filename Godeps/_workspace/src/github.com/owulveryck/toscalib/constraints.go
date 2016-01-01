package toscalib

type Value string

//Constraints is an array of ConstraintClause
type Constraints []ConstraintClause

// IsValid returns true if the Value is valid against the Constraints
func (c Constraints) IsValid(v Value) (bool, error) {
	return true, nil
}

// ConstraintClause definition as described in Appendix 5.2.
// This is a map where the index is a string that may have a value in
// {"equal","greater_than", ...} (see Appendix 5.2) a,s value is an interface
// for the definition.
// Example: ConstraintClause may be [ "greater_than": 3 ]
type ConstraintClause struct {
	Operator string
	Values   interface{}
}

// Evaluate the constraint and return a boolean
func (constraint *ConstraintClause) Evaluate(interface{}) bool { return true }

// UnmarshalYAML TODO: implement the Mashaler YAML interface for the constraint type
func (constraint *ConstraintClause) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var c map[string]interface{}
	err := unmarshal(&c)
	if err != nil {
		return err
	}
	var o string
	var v interface{}
	for op, val := range c {
		o = op
		v = val

	}
	constraint = &ConstraintClause{o, v}
	//*constraint = ConstraintClause(c)
	return nil
}
