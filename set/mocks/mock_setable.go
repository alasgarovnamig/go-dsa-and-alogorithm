package mocks

import (
	"github.com/alasgarovnamig/go-dsa-and-algorithm/set"
	"reflect"
)

// MockSetable is a mock implementation of the Setable interface,
// used for testing purposes.
type MockSetable struct {
	ID string
}

// NewMockSetable creates a new instance of MockSetable with a specific ID.
func NewMockSetable(id string) *MockSetable {
	return &MockSetable{ID: id}
}

// Hash returns the ID as the unique identifier for MockSetable.
func (m *MockSetable) Hash() string {
	return m.ID
}

// Equal compares MockSetable instances based on their ID.
// It returns false if the other object is nil or not of type *MockSetable.
func (m *MockSetable) Equal(other set.Setable) bool {
	// Check if `other` is nil or if it is a nil pointer within an interface.
	if other == nil || reflect.ValueOf(other).IsNil() {
		return false
	}

	// Check if `other` is of the exact type *MockSetable
	if reflect.TypeOf(other) != reflect.TypeOf(m) {
		return false
	}

	// Perform the type assertion since we know `other` is *MockSetable
	otherMock := other.(*MockSetable)
	return m.ID == otherMock.ID
}
