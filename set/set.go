package set

import "encoding/json"

// Set is a generic set data structure with JSON support
type Set[T comparable] struct {
	elements map[T]struct{}
}

// NewSet creates and returns a new Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{elements: make(map[T]struct{})}
}

// Add adds an element to the set
func (s *Set[T]) Add(value T) {
	s.elements[value] = struct{}{}
}

// Remove removes an element from the set
func (s *Set[T]) Remove(value T) {
	delete(s.elements, value)
}

// Contains checks if the set contains an element
func (s *Set[T]) Contains(value T) bool {
	_, exists := s.elements[value]
	return exists
}

// Size returns the size of the set
func (s *Set[T]) Size() int {
	return len(s.elements)
}

// Values returns all elements in the set as a slice
func (s *Set[T]) Values() []T {
	keys := make([]T, 0, len(s.elements))
	for key := range s.elements {
		keys = append(keys, key)
	}
	return keys
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.elements = make(map[T]struct{})
}

// MarshalJSON implements the json.Marshaler interface
func (s *Set[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *Set[T]) UnmarshalJSON(data []byte) error {
	var values []T
	if err := json.Unmarshal(data, &values); err != nil {
		return err
	}

	// Clear existing elements and add the new ones from JSON
	s.Clear()
	for _, value := range values {
		s.Add(value)
	}

	return nil
}
