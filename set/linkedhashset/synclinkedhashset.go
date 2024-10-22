package linkedhashset

import (
	"container/list"
	"fmt"
	"strings"
	"sync"
)

// SyncLinkedHashSet is a thread-safe version of LinkedHashSet.
type SyncLinkedHashSet[T comparable] struct {
	data  map[T]*list.Element
	order *list.List
	mu    sync.RWMutex
}

// NewSync initializes a new SyncLinkedHashSet.
func NewSync[T comparable]() *SyncLinkedHashSet[T] {
	return &SyncLinkedHashSet[T]{
		data:  make(map[T]*list.Element),
		order: list.New(),
	}
}

// Add inserts a value into the SyncLinkedHashSet.
func (s *SyncLinkedHashSet[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.data[value]; !exists {
		elem := s.order.PushBack(value)
		s.data[value] = elem
	}
}

// Remove deletes a value from the SyncLinkedHashSet.
func (s *SyncLinkedHashSet[T]) Remove(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if elem, exists := s.data[value]; exists {
		s.order.Remove(elem)
		delete(s.data, value)
	}
}

// Contains checks if a value exists in the SyncLinkedHashSet.
func (s *SyncLinkedHashSet[T]) Contains(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, exists := s.data[value]
	return exists
}

// Size returns the number of elements in the SyncLinkedHashSet.
func (s *SyncLinkedHashSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *SyncLinkedHashSet[T]) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Size() == 0
}
func (s *SyncLinkedHashSet[T]) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = make(map[T]*list.Element)
}
func (s *SyncLinkedHashSet[T]) Values() []T {
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
func (s *SyncLinkedHashSet[T]) ToString() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	str := "SyncLinkedHashSet : ["
	items := []string{}
	for k := range s.data {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	str += "]"
	return str
}
