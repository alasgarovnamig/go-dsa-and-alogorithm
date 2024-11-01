package hashset_test

import (
	"testing"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/hashset"
	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHashSet_BasicOperations(t *testing.T) {
	// HashSet oluştur
	hashSet := hashset.NewHashSet[*mocks.MockSetable]()

	// MockSetable nesneleri oluştur
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")
	item3 := mocks.NewMockSetable("item3")

	// Add metodunu test et
	hashSet.Add(item1, item2)
	assert.Equal(t, 2, hashSet.Size())
	assert.True(t, hashSet.Contains(item1))
	assert.True(t, hashSet.Contains(item2))
	assert.False(t, hashSet.Contains(item3))

	// Remove metodunu test et
	hashSet.Remove(item1)
	assert.False(t, hashSet.Contains(item1))
	assert.True(t, hashSet.Contains(item2))
	assert.Equal(t, 1, hashSet.Size())

	// Clear metodunu test et
	hashSet.Clear()
	assert.True(t, hashSet.IsEmpty())
	assert.Equal(t, 0, hashSet.Size())
}

func TestHashSet_ToString(t *testing.T) {
	hashSet := hashset.NewHashSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Set'e öğeler ekle
	hashSet.Add(item1, item2)

	// ToString metodunu test et
	str := hashSet.ToString()
	assert.Contains(t, str, "item1")
	assert.Contains(t, str, "item2")
}

func TestHashSet_ToSlice(t *testing.T) {
	hashSet := hashset.NewHashSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Set'e öğeler ekle
	hashSet.Add(item1, item2)

	// ToSlice metodunu test et
	slice := hashSet.ToSlice()
	assert.Len(t, slice, 2)
	assert.Contains(t, slice, item1)
	assert.Contains(t, slice, item2)
}
