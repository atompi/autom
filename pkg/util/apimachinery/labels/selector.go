package labels

import "github.com/atompi/autom/pkg/util/apimachinery/selection"

// Requirement contains values, a key, and an operator that relates the key and values.
// The zero value of Requirement is invalid.
// Requirement implements both set based match and exact match
// Requirement should be initialized via NewRequirement constructor for creating a valid Requirement.
// +k8s:deepcopy-gen=true
type Requirement struct {
	key      string
	operator selection.Operator
	// In huge majority of cases we have at most one value here.
	// It is generally faster to operate on a single-element slice
	// than on a single-element map, so we have a slice here.
	strValues []string
}

// Requirements is AND of all requirements.
type Requirements []Requirement

// Selector represents a label selector.
type Selector interface {
	// Matches returns true if this selector matches the given set of labels.
	Matches(Labels) bool

	// Empty returns true if this selector does not restrict the selection space.
	Empty() bool

	// String returns a human readable string that represents this selector.
	String() string

	// Add adds requirements to the Selector
	Add(r ...Requirement) Selector

	// Requirements converts this interface into Requirements to expose
	// more detailed selection information.
	// If there are querying parameters, it will return converted requirements and selectable=true.
	// If this selector doesn't want to select anything, it will return selectable=false.
	Requirements() (requirements Requirements, selectable bool)

	// Make a deep copy of the selector.
	DeepCopySelector() Selector

	// RequiresExactMatch allows a caller to introspect whether a given selector
	// requires a single specific label to be set, and if so returns the value it
	// requires.
	RequiresExactMatch(label string) (value string, found bool)
}
