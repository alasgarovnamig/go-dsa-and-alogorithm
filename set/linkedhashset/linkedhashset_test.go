package linkedhashset_test

import (
	"github.com/alasgarovnamig/go-dsa-and-alogorithm/set/linkedhashset"
	"testing"
)

func TestLinkedHashSet_Add(t *testing.T) {
	set := linkedhashset.New[int]()

	set.Add(1)
	set.Add(2)

	if !set.Contains(1) || !set.Contains(2) {
		t.Errorf("Expected set to contain 1 and 2")
	}
}

func TestLinkedHashSet_Remove(t *testing.T) {
	set := linkedhashset.New[int]()

	set.Add(1)
	set.Add(2)
	set.Remove(1)

	if set.Contains(1) {
		t.Errorf("Expected set to not contain 1")
	}

	if !set.Contains(2) {
		t.Errorf("Expected set to contain 2")
	}
}

func TestLinkedHashSet_Size(t *testing.T) {
	set := linkedhashset.New[int]()

	if set.Size() != 0 {
		t.Errorf("Expected size to be 0, got %d", set.Size())
	}

	set.Add(1)
	set.Add(2)

	if set.Size() != 2 {
		t.Errorf("Expected size to be 2, got %d", set.Size())
	}
}

func TestLinkedHashSet_IsEmpty(t *testing.T) {
	set := linkedhashset.New[int]()

	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty")
	}

	set.Add(1)

	if set.IsEmpty() {
		t.Errorf("Expected set to not be empty")
	}
}

func TestLinkedHashSet_Clear(t *testing.T) {
	set := linkedhashset.New[int]()

	set.Add(1)
	set.Add(2)
	set.Clear()

	if !set.IsEmpty() {
		t.Errorf("Expected set to be empty after Clear")
	}
}

func TestLinkedHashSet_Values(t *testing.T) {
	set := linkedhashset.New[int]()

	set.Add(1)
	set.Add(2)

	values := set.Values()
	expectedOrder := []int{1, 2}

	if len(values) != 2 {
		t.Errorf("Expected values length to be 2, got %d", len(values))
	}

	for i, v := range values {
		if v != expectedOrder[i] {
			t.Errorf("Unexpected value: %d, expected: %d", v, expectedOrder[i])
		}
	}
}

func TestLinkedHashSet_ToString(t *testing.T) {
	set := linkedhashset.New[int]()

	set.Add(1)
	set.Add(2)

	str := set.ToString()
	if str != "[1, 2]" {
		t.Errorf("Unexpected string representation: %s", str)
	}
}
