// Package toscalib implements the TOSCA syntax in its YAML version as described in
// http://docs.oasis-open.org/tosca/TOSCA-Simple-Profile-YAML/v1.0/csd03/TOSCA-Simple-Profile-YAML-v1.0-csd03.html
package toscalib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// This implements the type defined in Appendix A 2 of the definition file

// Version - The version have the following grammar:
// MajorVersion.MinorVersion[.FixVersion[.Qualifier[-BuildVersion]]]
// MajorVersion : is a required integer value greater than or equ al to 0 (zero)
// MinorVersion : is a required integer value greater than or equal to 0 (zero).
// FixVersion    : is a optional integer value greater than or equal to 0 (zero)
//Qualifier is an optional string that indicates a named, pre-release version of the associated code that has been derived from the version of the code identified by the combination major_version, minor_version and fix_version numbers
//BuildVersion is an optional integer value greater than or equal to 0 (zero) that can be used to further qualify different build versions of the code that has the same qualifer_string
type Version string

/*TODO
// GetMajor returns the major_version number
func (toscaVersion *Version) GetMajor() int {
	return 0
}
*/

/*TODO
// GetMinor returns the minor_version number
func (toscaVersion *Version) GetMinor() int {
	return 0
}
*/

/*TODO
// GetFixVersion returns the fix_version integer value
func (toscaVersion *Version) GetFixVersion() int {
	return 0
}
*/

/*TODO
// GetQualifier returns the named, pre-release version of the associated code that has been derived    from the version of the code identified by the combination major_version, minor_version and fix_version numbers
func (toscaVersion *Version) GetQualifier() string {
	return nil
}
*/

/*TODO
// GetBuildVersion returns an  integer value greater than or equal to 0 (zero) that can be used to further        qualify different build versions of the code that has the same qualifer_string
func (toscaVersion *Version) GetBuildVersion() int {
	return 0
}
*/

// UNBOUNDED: A.2.3 TOCSA range type
const UNBOUNDED uint64 = 9223372036854775807

// ToscaRange is defined in Appendix 2.3
// The range type can be used to define numeric ranges with a lower and upper boundary. For example, this allows for specifying a range of ports to be opened in a firewall
type ToscaRange interface{}

// ToscaList is defined is Appendix 2.4.
// The list type allows for specifying multiple values for a parameter of property.
// For example, if an application allows for being configured to listen on multiple ports, a list of ports could be configured using the list data type.
// Note that entries in a list for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by the entry_schema attribute of the respective property definition, attribute definitions, or input or output parameter definitions.
type ToscaList []interface{}

// ToscaMap type as described in appendix A.2.5
// The map type allows for specifying multiple values for a param eter of property as a map.
// In contrast to the list type, where each entry can only be addressed by its index in the list, entries in a map are named elements that can be addressed by their keys.i
// Note that entries in a map for one property or parameter must be of the same type.
// The type (for simple entries) or schema (for complex entries) is defined by
// the entry_schema attribute of the respective property definition, attribute definition, or input or output parameter definition
type ToscaMap map[interface{}]interface{}

// Size type as described in appendix A 2.6.4
type Size int64

// Frequency type as described in appendix A 2.6.6
type Frequency int64

// Scalar type as defined in Appendis 2.6.
// The scalar unit type can be used to define scalar values along with a unit from the list of recognized units
// Scalar type may be time.Duration, Size or Frequency
type Scalar struct {
	Value float64
	Unit  string
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
// Unmarshals a string of the form "scalar unit" into a Scalar, validating that scalar and unit are valid
func (s *Scalar) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var sString string
	err := unmarshal(&sString)
	if err != nil {
		return err
	}
	// Check if the s has two fields (one for the value, and the other one for the unit)
	ss := strings.Fields(sString)
	if len(ss) > 2 {
		return fmt.Errorf("Not a TOSCA scalar")
	}
	re := regexp.MustCompile("^([0-9.]+)[[:blank:]]*(B|kB|KiB|MB|MiB|GB|GiB|TB|TiB|d|h|m|s|ms|us|ns|Hz|kHz|MHz|GHz)$")
	res := re.FindStringSubmatch(sString)
	if err != nil || len(res) != 3 {
		return fmt.Errorf("Tosca type unkown")
	}
	val, err := strconv.ParseFloat(res[1], 64)
	if err != nil || len(res) != 3 {
		return fmt.Errorf("Not a number", res[1])
	}
	(*s).Value = val
	(*s).Unit = res[2]
	return nil
}

// Regex type used in the constraint definition (Appendix A 5.2.1)
type Regex interface{}
