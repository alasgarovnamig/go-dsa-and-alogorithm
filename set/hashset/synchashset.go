package hashset

import (
	"strings"
	"sync"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set"
)

// SyncHashSet is a thread-safe implementation of the HashSet.
// It uses a read-write mutex to allow concurrent access.
type SyncHashSet[T set.Setable] struct {
	elements map[string]T
	mu       sync.RWMutex
}

// NewSyncHashSet creates and returns a new instance of SyncHashSet.
func NewSyncHashSet[T set.Setable]() *SyncHashSet[T] {
	return &SyncHashSet[T]{
		elements: make(map[string]T),
	}
}

// Add inserts one or more elements into the SyncHashSet.
func (s *SyncHashSet[T]) Add(values ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, value := range values {
		key := value.Hash()
		s.elements[key] = value
	}
}

// Remove deletes one or more elements from the SyncHashSet.
func (s *SyncHashSet[T]) Remove(values ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, value := range values {
		key := value.Hash()
		delete(s.elements, key)
	}
}

// Contains checks if all specified elements are in the SyncHashSet.
func (s *SyncHashSet[T]) Contains(values ...T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, value := range values {
		key := value.Hash()
		if _, exists := s.elements[key]; !exists {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the SyncHashSet.
func (s *SyncHashSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.elements)
}

// IsEmpty checks if the SyncHashSet is empty.
func (s *SyncHashSet[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Clear removes all elements from the SyncHashSet.
func (s *SyncHashSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elements = make(map[string]T)
}

// ToString returns a string representation of the SyncHashSet.
func (s *SyncHashSet[T]) ToString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var sb strings.Builder
	sb.WriteString("SyncHashSet{")
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

// ToSlice returns a slice containing all elements in the SyncHashSet.
func (s *SyncHashSet[T]) ToSlice() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	slice := make([]T, 0, len(s.elements))
	for _, value := range s.elements {
		slice = append(slice, value)
	}
	return slice
}
