package hashset

import (
	"fmt"
	"testing"
)

// Equal function for integers
func intEqual(a, b int) bool {
	return a == b
}

// Hash function for integers
func intHash(v int) string {
	return fmt.Sprintf("%d", v)
}

// Equal function for strings
func stringEqual(a, b string) bool {
	return a == b
}

// Hash function for strings
func stringHash(v string) string {
	return v
}

func TestNewHashSet(t *testing.T) {
	// Initialize a new integer HashSet
	set := NewHashSet(intEqual, intHash, 1, 2, 3)

	// Check initial size
	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	// Check if elements exist
	if !set.Contains(1, 2, 3) {
		t.Errorf("Expected set to contain 1, 2, and 3")
	}

	// Check if an element doesn't exist
	if set.Contains(4) {
		t.Errorf("Expected set to not contain 4")
	}
}

func TestHashSet_Add(t *testing.T) {
	set := NewHashSet(intEqual, intHash)

	// Add elements
	set.Add(1, 2, 3)

	// Check size after adding elements
	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	// Check if elements exist
	if !set.Contains(1, 2, 3) {
		t.Errorf("Expected set to contain 1, 2, and 3")
	}

	// Add duplicate elements
	set.Add(2, 3)

	// Check size to ensure duplicates are not added
	if set.Size() != 3 {
		t.Errorf("Expected size to remain 3, got %d", set.Size())
	}
}

func TestHashSet_Remove(t *testing.T) {
	set := NewHashSet(intEqual, intHash, 1, 2, 3)

	// Remove elements
	set.Remove(2)

	// Check size after removal
	if set.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", set.Size())
	}

	// Check if element was removed
	if set.Contains(2) {
		t.Errorf("Expected set to not contain 2")
	}

	// Remove non-existing element
	set.Remove(4)

	// Ensure size doesn't change
	if set.Size() != 2 {
		t.Errorf("Expected size to remain 2, got %d", set.Size())
	}
}

func TestHashSet_Contains(t *testing.T) {
	set := NewHashSet(stringEqual, stringHash)

	// Add elements
	set.Add("a", "b", "c")

	// Check if elements exist
	if !set.Contains("a", "b", "c") {
		t.Errorf("Expected set to contain 'a', 'b', and 'c'")
	}

	// Check for non-existing element
	if set.Contains("d") {
		t.Errorf("Expected set to not contain 'd'")
	}
}

func TestHashSet_Size(t *testing.T) {
	set := NewHashSet(intEqual, intHash)

	// Check initial size
	if set.Size() != 0 {
		t.Errorf("Expected initial size to be 0, got %d", set.Size())
	}

	// Add elements
	set.Add(1, 2, 3)

	// Check size after adding elements
	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	// Remove an element
	set.Remove(2)

	// Check size after removal
	if set.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", set.Size())
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	set := NewHashSet(intEqual, intHash)

	// Check if set is empty initially
	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	// Add elements
	set.Add(1)

	// Check if set is not empty
	if set.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}

	// Clear the set
	set.Clear()

	// Check if set is empty again
	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty after clearing")
	}
}

func TestHashSet_Clear(t *testing.T) {
	set := NewHashSet(intEqual, intHash, 1, 2, 3)

	// Clear the set
	set.Clear()

	// Check if set is empty
	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty after clearing")
	}

	// Check size after clearing
	if set.Size() != 0 {
		t.Errorf("Expected size to be 0 after clearing, got %d", set.Size())
	}
}

func TestHashSet_Values(t *testing.T) {
	set := NewHashSet(intEqual, intHash, 1, 2, 3)

	// Get values
	values := set.Values()

	// Create a map for validation
	expected := map[int]bool{1: true, 2: true, 3: true}

	// Check length of values slice
	if len(values) != 3 {
		t.Errorf("Expected values length to be 3, got %d", len(values))
	}

	// Check if all values are present
	for _, v := range values {
		if !expected[v] {
			t.Errorf("Unexpected value: %d", v)
		}
	}
}

func TestHashSet_ToString(t *testing.T) {
	set := NewHashSet(intEqual, intHash, 1, 2, 3)

	// Get string representation
	str := set.ToString()

	// Check if string representation is correct
	if str != "[1, 2, 3]" && str != "[1, 3, 2]" && str != "[2, 1, 3]" && str != "[2, 3, 1]" && str != "[3, 1, 2]" && str != "[3, 2, 1]" {
		t.Errorf("Unexpected string representation: %s", str)
	}
}

func TestHashSet_Add_Remove_Multiple(t *testing.T) {
	set := NewHashSet(intEqual, intHash)

	// Add multiple elements
	set.Add(1, 2, 3, 4, 5)

	// Check size after adding multiple elements
	if set.Size() != 5 {
		t.Errorf("Expected size to be 5, got %d", set.Size())
	}

	// Remove multiple elements
	set.Remove(1, 2)

	// Check size after removing multiple elements
	if set.Size() != 3 {
		t.Errorf("Expected size to be 3, got %d", set.Size())
	}

	// Check if removed elements are not present
	if set.Contains(1, 2) {
		t.Errorf("Expected set to not contain 1 and 2")
	}

	// Check if remaining elements are present
	if !set.Contains(3, 4, 5) {
		t.Errorf("Expected set to contain 3, 4, and 5")
	}
}
