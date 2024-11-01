package hashset

import (
	"strings"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set"
)

// HashSet implements the Set interface using a map for uniqueness.
// It maintains unique elements in no particular order.
type HashSet[T set.Setable] struct {
	elements map[string]T
}

// NewHashSet creates and returns a new instance of HashSet.
func NewHashSet[T set.Setable]() *HashSet[T] {
	return &HashSet[T]{
		elements: make(map[string]T),
	}
}

// Add inserts one or more elements into the HashSet.
// Duplicate elements (based on Hash) are ignored.
func (s *HashSet[T]) Add(values ...T) {
	for _, value := range values {
		key := value.Hash()
		s.elements[key] = value // Overwrite if already exists
	}
}

// Remove deletes one or more elements from the HashSet.
func (s *HashSet[T]) Remove(values ...T) {
	for _, value := range values {
		key := value.Hash()
		delete(s.elements, key)
	}
}

// Contains checks if all specified elements are present in the HashSet.
func (s *HashSet[T]) Contains(values ...T) bool {
	for _, value := range values {
		key := value.Hash()
		if _, exists := s.elements[key]; !exists {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the HashSet.
func (s *HashSet[T]) Size() int {
	return len(s.elements)
}

// IsEmpty checks if the HashSet has no elements.
func (s *HashSet[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

// Clear removes all elements from the HashSet.
func (s *HashSet[T]) Clear() {
	s.elements = make(map[string]T)
}

// ToString returns a string representation of the HashSet.
func (s *HashSet[T]) ToString() string {
	var sb strings.Builder
	sb.WriteString("HashSet{")
	first := true
	for _, value := range s.elements {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(value.Hash())
		first = false
	}
	sb.WriteString("}")
	return sb.String()
}

// ToSlice returns a slice containing all elements in the HashSet.
func (s *HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for _, value := range s.elements {
		slice = append(slice, value)
	}
	return slice
}
