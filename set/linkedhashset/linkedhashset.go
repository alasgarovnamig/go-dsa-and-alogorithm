package linkedhashset

import (
	"container/list"
	"fmt"
	"strings"
)

// LinkedHashSet maintains insertion order of elements.
type LinkedHashSet[T comparable] struct {
	data  map[T]*list.Element
	order *list.List
}

// New initializes a new LinkedHashSet.
func New[T comparable]() *LinkedHashSet[T] {
	return &LinkedHashSet[T]{
		data:  make(map[T]*list.Element),
		order: list.New(),
	}
}

// Add inserts a value into the LinkedHashSet.
func (s *LinkedHashSet[T]) Add(value T) {
	if _, exists := s.data[value]; !exists {
		elem := s.order.PushBack(value)
		s.data[value] = elem
	}
}

// Remove deletes a value from the LinkedHashSet.
func (s *LinkedHashSet[T]) Remove(value T) {
	if elem, exists := s.data[value]; exists {
		s.order.Remove(elem)
		delete(s.data, value)
	}
}

// Contains checks if a value exists in the LinkedHashSet.
func (s *LinkedHashSet[T]) Contains(value T) bool {
	_, exists := s.data[value]
	return exists
}

// Size returns the number of elements in the LinkedHashSet.
func (s *LinkedHashSet[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks if the LinkedHashSet is empty.
func (s *LinkedHashSet[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Clear removes all elements from the LinkedHashSet.
func (s *LinkedHashSet[T]) Clear() {
	s.data = make(map[T]*list.Element)
	s.order.Init()
}

// Values returns a slice of all elements in the LinkedHashSet.
func (s *LinkedHashSet[T]) Values() []T {
	values := make([]T, 0, len(s.data))
	for elem := s.order.Front(); elem != nil; elem = elem.Next() {
		values = append(values, elem.Value.(T))
	}
	return values
}

// ToString returns a string representation of the LinkedHashSet.
func (s *LinkedHashSet[T]) ToString() string {
	var sb strings.Builder
	sb.WriteString("[")
	first := true
	for elem := s.order.Front(); elem != nil; elem = elem.Next() {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", elem.Value))
		first = false
	}
	sb.WriteString("]")
	return sb.String()
}
