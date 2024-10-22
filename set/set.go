package set

// Set defines the basic operations for a set data structure.
type Set[T comparable] interface {
	Add(value T)           // Adds an element to the set
	Remove(value T)        // Removes an element from the set
	Contains(value T) bool // Checks if an element exists in the set
	Size() int             // Returns the size of the set
	IsEmpty() bool
	Clear()
	Values() []T
	ToString() string
}
