package treeset

import "sync"

// SyncTreeSet is a thread-safe version of TreeSet with a custom comparator.
type SyncTreeSet[T any] struct {
	root       *TreeNode[T]
	size       int
	comparator Comparator[T]
	mu         sync.RWMutex
}

// NewSync initializes a new SyncTreeSet with a comparator function.
func NewSync[T any](comparator Comparator[T]) *SyncTreeSet[T] {
	return &SyncTreeSet[T]{comparator: comparator}
}

// Add inserts a value into the SyncTreeSet.
func (s *SyncTreeSet[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.root == nil {
		s.root = &TreeNode[T]{value: value}
	} else {
		s.root = s.addRecursive(s.root, value)
	}
	s.size++
}

func (s *SyncTreeSet[T]) addRecursive(node *TreeNode[T], value T) *TreeNode[T] {
	if node == nil {
		return &TreeNode[T]{value: value}
	}

	if s.comparator(value, node.value) < 0 {
		node.left = s.addRecursive(node.left, value)
	} else if s.comparator(value, node.value) > 0 {
		node.right = s.addRecursive(node.right, value)
	}
	return node
}

// Contains checks if a value exists in the SyncTreeSet.
func (s *SyncTreeSet[T]) Contains(value T) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.containsRecursive(s.root, value)
}

func (s *SyncTreeSet[T]) containsRecursive(node *TreeNode[T], value T) bool {
	if node == nil {
		return false
	}

	if s.comparator(value, node.value) == 0 {
		return true
	} else if s.comparator(value, node.value) < 0 {
		return s.containsRecursive(node.left, value)
	}
	return s.containsRecursive(node.right, value)
}

// Size returns the number of elements in the SyncTreeSet.
func (s *SyncTreeSet[T]) Size() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.size
}
