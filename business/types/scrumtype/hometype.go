// Package scrumtype represents the scrum type in the system.
package scrumtype

import "fmt"

// The set of types that can be used.
var (
	Single = newType("SINGLE FAMILY")
	Condo  = newType("CONDO")
)

// =============================================================================

// Set of known housing types.
var scrumTypes = make(map[string]ScrumType)

// ScrumType represents a type in the system.
type ScrumType struct {
	value string
}

func newType(scrumType string) ScrumType {
	ht := ScrumType{scrumType}
	scrumTypes[scrumType] = ht
	return ht
}

// String returns the name of the type.
func (ht ScrumType) String() string {
	return ht.value
}

// Equal provides support for the go-cmp package and testing.
func (ht ScrumType) Equal(ht2 ScrumType) bool {
	return ht.value == ht2.value
}

// MarshalText provides support for logging and any marshal needs.
func (ht ScrumType) MarshalText() ([]byte, error) {
	return []byte(ht.value), nil
}

// =============================================================================

// Parse parses the string value and returns a scrum type if one exists.
func Parse(value string) (ScrumType, error) {
	typ, exists := scrumTypes[value]
	if !exists {
		return ScrumType{}, fmt.Errorf("invalid scrum type %q", value)
	}

	return typ, nil
}

// MustParse parses the string value and returns a scrum type if one exists. If
// an error occurs the function panics.
func MustParse(value string) ScrumType {
	typ, err := Parse(value)
	if err != nil {
		panic(err)
	}

	return typ
}
