package set

// Setable defines an interface for elements that can be used in a set.
// It requires two methods: Hash for generating a unique identifier and
// Equal for checking equality with another Setable instance.
type Setable interface {
	// Hash returns a unique identifier for the element.
	// This identifier is used for determining element uniqueness in the set.
	Hash() string

	// Equal checks if the current element is equal to another element
	// based on certain criteria defined by the element itself.
	Equal(other Setable) bool
}
