package set

import (
	"container/list"
	"encoding/json"
)

// LinkedSet is a set that maintains the insertion order of elements
type LinkedSet[T any] struct {
	elements map[interface{}]*list.Element
	order    *list.List
}

// NewLinkedSet creates a new LinkedSet
func NewLinkedSet[T any]() *LinkedSet[T] {
	return &LinkedSet[T]{
		elements: make(map[interface{}]*list.Element),
		order:    list.New(),
	}
}

// Add adds an element to the set, maintaining the insertion order
func (s *LinkedSet[T]) Add(value T) {
	key := getKey(value)
	if _, exists := s.elements[key]; !exists {
		elem := s.order.PushBack(value)
		s.elements[key] = elem
	}
}

// Remove removes an element from the set
func (s *LinkedSet[T]) Remove(value T) {
	key := getKey(value)
	if elem, exists := s.elements[key]; exists {
		s.order.Remove(elem)
		delete(s.elements, key)
	}
}

// Contains checks if the set contains an element
func (s *LinkedSet[T]) Contains(value T) bool {
	key := getKey(value)
	_, exists := s.elements[key]
	return exists
}

// Size returns the size of the set
func (s *LinkedSet[T]) Size() int {
	return len(s.elements)
}

// Values returns all elements in the set as a slice, in insertion order
func (s *LinkedSet[T]) Values() []T {
	values := make([]T, 0, len(s.elements))
	for elem := s.order.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(T))
	}
	return values
}

// Clear removes all elements from the set
func (s *LinkedSet[T]) Clear() {
	s.elements = make(map[interface{}]*list.Element)
	s.order.Init()
}

// MarshalJSON implements the json.Marshaler interface
func (s *LinkedSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *LinkedSet[T]) UnmarshalJSON(data []byte) error {
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
