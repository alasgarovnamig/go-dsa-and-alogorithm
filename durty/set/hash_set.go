package set

import (
	"encoding/json"
	"reflect"
)

// Set is a generic set data structure with optional HashCode method support
type HashSet[T any] struct {
	elements map[interface{}]T
}

// NewSet creates and returns a new Set
func NewHashSet[T any]() *HashSet[T] {
	return &HashSet[T]{elements: make(map[interface{}]T)}
}

// Add adds an element to the set, using HashCode() if available, or default comparison
func (s *HashSet[T]) Add(value T) {
	key := getKey(value)
	s.elements[key] = value
}

// Remove removes an element from the set
func (s *HashSet[T]) Remove(value T) {
	key := getKey(value)
	delete(s.elements, key)
}

// Contains checks if the set contains an element
func (s *HashSet[T]) Contains(value T) bool {
	key := getKey(value)
	_, exists := s.elements[key]
	return exists
}

// Size returns the size of the set
func (s *HashSet[T]) Size() int {
	return len(s.elements)
}

// Values returns all elements in the set as a slice
func (s *HashSet[T]) Values() []T {
	values := make([]T, 0, len(s.elements))
	for _, value := range s.elements {
		values = append(values, value)
	}
	return values
}

// Clear removes all elements from the set
func (s *HashSet[T]) Clear() {
	s.elements = make(map[interface{}]T)
}

// getKey generates a unique key based on the element, using HashCode() if available
func getKey[T any](value T) interface{} {
	v := reflect.ValueOf(value)

	// Check if the struct has a method named "HashCode"
	if v.Kind() == reflect.Struct {
		method := v.MethodByName("HashCode")
		if method.IsValid() && method.Type().NumIn() == 0 && method.Type().NumOut() == 1 {
			// HashCode method exists, call it
			return method.Call(nil)[0].Interface()
		}
	}

	// Fallback: use the value itself for comparison
	return v.Interface()
}

// MarshalJSON implements the json.Marshaler interface
func (s *HashSet[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values())
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (s *HashSet[T]) UnmarshalJSON(data []byte) error {
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
