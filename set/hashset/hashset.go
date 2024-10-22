package hashset

import (
	"fmt"
	"strings"
)

// HashSet is a simple set implemented using a map.
type HashSet[T comparable] struct {
	data map[T]struct{}
}

// NewHashSet initializes a new HashSet.
func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{data: make(map[T]struct{})}
}

// Add inserts a value into the HashSet.
func (s *HashSet[T]) Add(value T) {
	s.data[value] = struct{}{}
}

// Remove deletes a value from the HashSet.
func (s *HashSet[T]) Remove(value T) {
	delete(s.data, value)
}

// Contains checks if a value exists in the HashSet.
func (s *HashSet[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Size returns the number of elements in the HashSet.
func (s *HashSet[T]) Size() int {
	return len(s.data)
}

func (s *HashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}
func (s *HashSet[T]) Clear() {
	s.data = make(map[T]struct{})
}
func (s *HashSet[T]) Values() []T {
	values := make([]T, s.Size())
	count := 0
	for item := range s.data {
		values[count] = item
		count++
	}
	return values
}
func (s *HashSet[T]) ToString() string {
	str := "HashSet:["
	items := []string{}
	for k := range s.data {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	str += "]"
	return str
}
