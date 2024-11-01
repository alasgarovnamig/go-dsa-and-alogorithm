package hashset_test

import (
	"sync"
	"testing"

	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/hashset"
	"github.com/alasgarovnamig/go-dsa-and-algorithm/set/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSyncHashSet_BasicOperations(t *testing.T) {
	// SyncHashSet oluştur
	syncHashSet := hashset.NewSyncHashSet[*mocks.MockSetable]()

	// MockSetable nesneleri oluştur
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")
	item3 := mocks.NewMockSetable("item3")

	// Add metodunu test et
	syncHashSet.Add(item1, item2)
	assert.Equal(t, 2, syncHashSet.Size())
	assert.True(t, syncHashSet.Contains(item1))
	assert.True(t, syncHashSet.Contains(item2))
	assert.False(t, syncHashSet.Contains(item3))

	// Remove metodunu test et
	syncHashSet.Remove(item1)
	assert.False(t, syncHashSet.Contains(item1))
	assert.True(t, syncHashSet.Contains(item2))
	assert.Equal(t, 1, syncHashSet.Size())

	// Clear metodunu test et
	syncHashSet.Clear()
	assert.True(t, syncHashSet.IsEmpty())
	assert.Equal(t, 0, syncHashSet.Size())
}

func TestSyncHashSet_ToString(t *testing.T) {
	syncHashSet := hashset.NewSyncHashSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Set'e öğeler ekle
	syncHashSet.Add(item1, item2)

	// ToString metodunu test et
	str := syncHashSet.ToString()
	assert.Contains(t, str, "item1")
	assert.Contains(t, str, "item2")
}

func TestSyncHashSet_ToSlice(t *testing.T) {
	syncHashSet := hashset.NewSyncHashSet[*mocks.MockSetable]()
	item1 := mocks.NewMockSetable("item1")
	item2 := mocks.NewMockSetable("item2")

	// Set'e öğeler ekle
	syncHashSet.Add(item1, item2)

	// ToSlice metodunu test et
	slice := syncHashSet.ToSlice()
	assert.Len(t, slice, 2)
	assert.Contains(t, slice, item1)
	assert.Contains(t, slice, item2)
}

func TestSyncHashSet_ConcurrentAccess(t *testing.T) {
	syncHashSet := hashset.NewSyncHashSet[*mocks.MockSetable]()
	item := mocks.NewMockSetable("concurrentItem")

	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine 1: 100 kez aynı öğeyi ekler
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			syncHashSet.Add(item)
		}
	}()

	// Goroutine 2: 100 kez aynı öğeyi kontrol eder
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			_ = syncHashSet.Contains(item)
		}
	}()

	// Goroutineler tamamlanana kadar bekleyin
	wg.Wait()

	// Set'in boyutunu kontrol et
	assert.Equal(t, 1, syncHashSet.Size())
	assert.True(t, syncHashSet.Contains(item))
}
