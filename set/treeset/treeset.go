package treeset

// Comparator is a function type that defines the ordering of elements.
// It returns a negative value if a < b, zero if a == b, and a positive value if a > b.
type Comparator[T any] func(a, b T) int

// TreeSet is a sorted set implemented using a binary search tree.
type TreeSet[T any] struct {
	root       *TreeNode[T]
	size       int
	comparator Comparator[T]
}

// New initializes a new TreeSet with a custom comparator function.
func New[T any](comparator Comparator[T]) *TreeSet[T] {
	return &TreeSet[T]{comparator: comparator}
}

// Add inserts a value into the TreeSet.
func (s *TreeSet[T]) Add(value T) {
	if s.root == nil {
		s.root = &TreeNode[T]{value: value}
	} else {
		s.root = s.addRecursive(s.root, value)
	}
	s.size++
}

func (s *TreeSet[T]) addRecursive(node *TreeNode[T], value T) *TreeNode[T] {
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

// Contains checks if a value exists in the TreeSet.
func (s *TreeSet[T]) Contains(value T) bool {
	return s.containsRecursive(s.root, value)
}

func (s *TreeSet[T]) containsRecursive(node *TreeNode[T], value T) bool {
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

// Size returns the number of elements in the TreeSet.
func (s *TreeSet[T]) Size() int {
	return s.size
}
