package set

// Set defines the basic operations for a set data structure.
// It requires elements to implement the Setable interface for uniqueness checks.
type Set[T Setable] interface {
	// Add inserts one or more elements into the set.
	// Duplicate elements are ignored.
	Add(value ...T)

	// Remove deletes one or more elements from the set.
	Remove(value ...T)

	// Contains checks if all specified elements are in the set.
	Contains(value ...T) bool

	// Size returns the number of elements in the set.
	Size() int

	// IsEmpty checks if the set is empty.
	IsEmpty() bool

	// Clear removes all elements from the set.
	Clear()

	// ToString returns a string representation of the set.
	ToString() string

	// ToSlice returns a slice containing all elements in the set.
	ToSlice() []T
}
