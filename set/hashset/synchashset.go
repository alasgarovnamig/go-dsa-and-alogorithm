package hashset

import (
	"fmt"
	"strings"
	"sync"
)

// SyncHashSet is a concurrent set implemented using a map and mutex.
type SyncHashSet[T any] struct {
	data    map[string]T
	equalFn func(T, T) bool
	hashFn  func(T) string
	mu      sync.RWMutex
}

// NewSyncHashSet initializes a new SyncHashSet.
func NewSyncHashSet[T any](equalFn func(T, T) bool, hashFn func(T) string, data ...T) *SyncHashSet[T] {
	hashSet := &SyncHashSet[T]{
		data:    make(map[string]T),
		equalFn: equalFn,
		hashFn:  hashFn,
	}
	if len(data) > 0 {
		hashSet.Add(data...)
	}
	return hashSet
}

// Add inserts a value into the SyncHashSet.
func (s *SyncHashSet[T]) Add(data ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range data {
		hash := s.hashFn(v)
		s.data[hash] = v
	}
}

// Remove deletes a value from the SyncHashSet.
func (s *SyncHashSet[T]) Remove(data ...T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, v := range data {
		hash := s.hashFn(v)
		delete(s.data, hash)
	}
}

// Contains checks if a value exists in the SyncHashSet.
func (s *SyncHashSet[T]) Contains(data ...T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, v := range data {
		hash := s.hashFn(v)
		if _, contains := s.data[hash]; !contains {
			return false
		}
	}
	return true
}

// Size returns the number of elements in the SyncHashSet.
func (s *SyncHashSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// IsEmpty checks if the SyncHashSet is empty.
func (s *SyncHashSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data) == 0
}

// Clear removes all elements from the SyncHashSet.
func (s *SyncHashSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[string]T)
}

// Values returns a slice of all elements in the SyncHashSet.
func (s *SyncHashSet[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()

	values := make([]T, 0, len(s.data))
	for _, value := range s.data {
		values = append(values, value)
	}
	return values
}

// ToString returns a string representation of the SyncHashSet.
func (s *SyncHashSet[T]) ToString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()

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
