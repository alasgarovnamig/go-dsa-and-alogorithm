package mocks

import (
	"github.com/alasgarovnamig/go-dsa-and-algorithm/set"
	"strings"
)

// MockSet is a manually created mock implementation of the Set interface.
// It uses a map to store elements for uniqueness and simulates Set behavior.
type MockSet[T set.Setable] struct {
	items map[string]T
}

// NewMockSet creates a new instance of MockSet.
func NewMockSet[T set.Setable]() *MockSet[T] {
	return &MockSet[T]{items: make(map[string]T)}
}

// Add adds elements to the mock set.
func (m *MockSet[T]) Add(values ...T) {
	for _, value := range values {
		m.items[value.Hash()] = value
	}
}

// Remove removes elements from the mock set.
func (m *MockSet[T]) Remove(values ...T) {
	for _, value := range values {
		delete(m.items, value.Hash())
	}
}

// Contains checks if elements exist in the mock set.
func (m *MockSet[T]) Contains(values ...T) bool {
	for _, value := range values {
		if _, exists := m.items[value.Hash()]; !exists {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the mock set.
func (m *MockSet[T]) Size() int {
	return len(m.items)
}

// IsEmpty checks if the mock set is empty.
func (m *MockSet[T]) IsEmpty() bool {
	return len(m.items) == 0
}

// Clear removes all elements from the mock set.
func (m *MockSet[T]) Clear() {
	m.items = make(map[string]T)
}

// ToString returns a string representation of the mock set.
func (m *MockSet[T]) ToString() string {
	var sb strings.Builder
	sb.WriteString("MockSet{")
	first := true
	for _, value := range m.items {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(value.Hash())
		first = false
	}
	sb.WriteString("}")
	return sb.String()
}

// ToSlice returns a slice containing all elements in the mock set.
func (m *MockSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(m.items))
	for _, value := range m.items {
		slice = append(slice, value)
	}
	return slice
}
