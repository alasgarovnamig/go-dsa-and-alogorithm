package treeset

// TreeNode represents a node in the binary search tree.
type TreeNode[T any] struct {
	value T
	left  *TreeNode[T]
	right *TreeNode[T]
}
