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

func (s *LinkedHashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}
func (s *LinkedHashSet[T]) Clear() {
	s.data = make(map[T]*list.Element)
}
func (s *LinkedHashSet[T]) Values() []T {
	values := make([]T, s.Size())
	count := 0
	for item := range s.data {
		values[count] = item
		count++
	}
	return values
}
func (s *LinkedHashSet[T]) ToString() string {
	str := "LinkedHashSet:["
	items := []string{}
	for k := range s.data {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	str += "]"
	return str
}
