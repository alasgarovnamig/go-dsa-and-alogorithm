package hashset

import (
	"fmt"
	"strings"
	"sync"
)

// SyncHashSet is a concurrent version of the HashSet.
type SyncHashSet[T comparable] struct {
	data map[T]struct{}
	mu   sync.RWMutex
}

// NewSync initializes a new SyncHashSet.
func NewSync[T comparable]() *SyncHashSet[T] {
	return &SyncHashSet[T]{data: make(map[T]struct{})}
}

// Add inserts a value into the SyncHashSet.
func (s *SyncHashSet[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[value] = struct{}{}
}

// Remove deletes a value from the SyncHashSet.
func (s *SyncHashSet[T]) Remove(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, value)
}

// Contains checks if a value exists in the SyncHashSet.
func (s *SyncHashSet[T]) Contains(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.data[value]
	return exists
}

// Size returns the number of elements in the SyncHashSet.
func (s *SyncHashSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *SyncHashSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Size() == 0
}
func (s *SyncHashSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[T]struct{})
}
func (s *SyncHashSet[T]) Values() []T {
	s.mu.RLock()
	defer s.mu.RUnlock()
	values := make([]T, s.Size())
	count := 0
	for item := range s.data {
		values[count] = item
		count++
	}
	return values
}
func (s *SyncHashSet[T]) ToString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	str := "SyncHashSet:["
	items := []string{}
	for k := range s.data {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	str += "]"
	return str
}
