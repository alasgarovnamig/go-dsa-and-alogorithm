package mocks_test

import (
	"testing"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set"
	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/mocks"
	"github.com/stretchr/testify/assert"
)

func TestMockSetable_Hash(t *testing.T) {
	// Create a MockSetable instance
	mockItem := mocks.NewMockSetable("uniqueID")

	// Test Hash method
	assert.Equal(t, "uniqueID", mockItem.Hash(), "Hash should return the ID as unique identifier")
}

func TestMockSetable_Equal_SameID(t *testing.T) {
	// Create two MockSetable instances with the same ID
	mockItem1 := mocks.NewMockSetable("item1")
	mockItem2 := mocks.NewMockSetable("item1")

	// Test Equal method
	assert.True(t, mockItem1.Equal(mockItem2), "Items with the same ID should be equal")
}

func TestMockSetable_Equal_DifferentID(t *testing.T) {
	// Create two MockSetable instances with different IDs
	mockItem1 := mocks.NewMockSetable("item1")
	mockItem2 := mocks.NewMockSetable("item2")

	// Test Equal method
	assert.False(t, mockItem1.Equal(mockItem2), "Items with different IDs should not be equal")
}

func TestMockSetable_Equal_NilComparison(t *testing.T) {
	// Create a MockSetable instance and compare it with nil
	mockItem := mocks.NewMockSetable("item1")
	var nilItem *mocks.MockSetable = nil

	// Test Equal method with nil comparison
	assert.False(t, mockItem.Equal(nilItem), "Comparing with nil should return false")
}

func TestMockSetable_Equal_DifferentType(t *testing.T) {
	// Create a MockSetable instance and compare it with a different type
	mockItem := mocks.NewMockSetable("item1")

	// Pass a non-MockSetable instance to Equal
	var otherType set.Setable = mocks.NewMockSetable("otherItem")

	// Test Equal method with a different type
	assert.False(t, mockItem.Equal(otherType), "Comparing with a different Setable type should return false")
}
