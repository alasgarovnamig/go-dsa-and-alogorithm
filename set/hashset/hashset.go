package hashset

import (
	"fmt"
	"strings"
)

// HashSet is a simple set implemented using a map.
type HashSet[T any] struct {
	data    map[string]T
	equalFn func(T, T) bool
	hashFn  func(T) string
}

// New initializes a new HashSet.
func NewHashSet[T any](equalFn func(T, T) bool, hashFn func(T) string, data ...T) *HashSet[T] {
	hashSet := &HashSet[T]{
		data:    make(map[string]T),
		equalFn: equalFn,
		hashFn:  hashFn,
	}
	if len(data) > 0 {
		hashSet.Add(data...)
	}
	return hashSet
}

// Add inserts a value into the HashSet.
func (s *HashSet[T]) Add(data ...T) {
	for _, v := range data {
		hash := s.hashFn(v)
		s.data[hash] = v
	}
}

// Remove deletes a value from the HashSet.
func (s *HashSet[T]) Remove(data ...T) {
	for _, v := range data {
		hash := s.hashFn(v)
		delete(s.data, hash)
	}
}

// Contains checks if a value exists in the HashSet.
func (s *HashSet[T]) Contains(data ...T) bool {
	for _, v := range data {
		hash := s.hashFn(v)
		if _, contains := s.data[hash]; !contains {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the HashSet.
func (s *HashSet[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks if the HashSet is empty.
func (s *HashSet[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the HashSet.
func (s *HashSet[T]) Clear() {
	s.data = make(map[string]T)
}

// Values returns a slice of all elements in the HashSet.
func (s *HashSet[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for _, value := range s.data {
		values = append(values, value)
	}
	return values
}

// ToString returns a string representation of the HashSet.
func (s *HashSet[T]) ToString() string {
	var sb strings.Builder
	sb.WriteString("[")
	first := true
	for _, value := range s.Values() {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", value))
		first = false
	}
	sb.WriteString("]")
	return sb.String()

}
