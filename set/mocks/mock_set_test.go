package mocks_test

import (
	"testing"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockSet_Add(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Add metodunu test et
	mockSet.Add(item1, item2)
	assert.Equal(t, 2, mockSet.Size(), "Size should be 2 after adding two items")
	assert.True(t, mockSet.Contains(item1), "Set should contain item1")
	assert.True(t, mockSet.Contains(item2), "Set should contain item2")
}

func TestMockSet_Remove(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Öğeleri ekle ve sonra Remove metodunu test et
	mockSet.Add(item1, item2)
	mockSet.Remove(item1)
	assert.Equal(t, 1, mockSet.Size(), "Size should be 1 after removing one item")
	assert.False(t, mockSet.Contains(item1), "Set should not contain item1 after removal")
	assert.True(t, mockSet.Contains(item2), "Set should still contain item2")
}

func TestMockSet_Contains(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Contains metodunu test et
	mockSet.Add(item1)
	assert.True(t, mockSet.Contains(item1), "Set should contain item1")
	assert.False(t, mockSet.Contains(item2), "Set should not contain item2")
}

func TestMockSet_SizeAndIsEmpty(t *testing.T) {
	// MockSet oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	assert.True(t, mockSet.IsEmpty(), "Set should be empty initially")

	// Öğeleri ekle
	item1 := mocks.NewMockSetable("item1")
	mockSet.Add(item1)
	assert.Equal(t, 1, mockSet.Size(), "Size should be 1 after adding one item")
	assert.False(t, mockSet.IsEmpty(), "Set should not be empty after adding items")
}

func TestMockSet_Clear(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Öğeleri ekle ve sonra Clear metodunu test et
	mockSet.Add(item1, item2)
	mockSet.Clear()
	assert.Equal(t, 0, mockSet.Size(), "Size should be 0 after clearing the set")
	assert.True(t, mockSet.IsEmpty(), "Set should be empty after clearing")
}

func TestMockSet_ToSlice(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Öğeleri ekle ve ToSlice metodunu test et
	mockSet.Add(item1, item2)
	slice := mockSet.ToSlice()
	assert.Len(t, slice, 2, "Slice length should match the set size")
	assert.Contains(t, slice, item1, "Slice should contain item1")
	assert.Contains(t, slice, item2, "Slice should contain item2")
}

func TestMockSet_ToString(t *testing.T) {
	// MockSet ve MockSetable nesneleri oluştur
	mockSet := mocks.NewMockSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Öğeleri ekle ve ToString metodunu test et
	mockSet.Add(item1, item2)
	str := mockSet.ToString()
	assert.Contains(t, str, "item1", "String representation should contain item1")
	assert.Contains(t, str, "item2", "String representation should contain item2")
}
