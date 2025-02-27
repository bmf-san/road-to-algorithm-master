package hash_table

import (
	"fmt"
	"testing"
)

func TestSwissTableBasicOperations(t *testing.T) {
	// Create a new hash table
	st := NewSwissTable()

	// Test initial state
	if st.Size() != 0 {
		t.Errorf("Expected size 0, got %d", st.Size())
	}
	if st.Capacity() != minCapacity {
		t.Errorf("Expected capacity %d, got %d", minCapacity, st.Capacity())
	}

	// Test Insert
	if !st.Insert(1, "one") {
		t.Error("Insert failed")
	}
	if !st.Insert(2, "two") {
		t.Error("Insert failed")
	}
	if st.Size() != 2 {
		t.Errorf("Expected size 2, got %d", st.Size())
	}

	// Test Search
	if val, found := st.Search(1); !found || val != "one" {
		t.Error("Search failed for key 1")
	}
	if val, found := st.Search(2); !found || val != "two" {
		t.Error("Search failed for key 2")
	}
	if _, found := st.Search(3); found {
		t.Error("Search should not find non-existent key")
	}

	// Test collision handling
	if !st.Insert(17, "seventeen") { // 17 % 16 = 1, should handle collision
		t.Error("Insert failed on collision")
	}
	if val, found := st.Search(17); !found || val != "seventeen" {
		t.Error("Search failed for collided key")
	}

	// Test Delete
	if !st.Delete(1) {
		t.Error("Delete failed")
	}
	if _, found := st.Search(1); found {
		t.Error("Search found deleted key")
	}
	if st.Size() != 2 {
		t.Errorf("Expected size 2 after delete, got %d", st.Size())
	}
}

func TestSwissTableUpdates(t *testing.T) {
	st := NewSwissTable()

	// Test updating existing key
	st.Insert(1, "one")
	st.Insert(1, "ONE")
	if val, found := st.Search(1); !found || val != "ONE" {
		t.Error("Update existing key failed")
	}
	if st.Size() != 1 {
		t.Errorf("Size should be 1 after update, got %d", st.Size())
	}
}

func TestSwissTableEmptyOperations(t *testing.T) {
	st := NewSwissTable()

	// Test operations on empty table
	if _, found := st.Search(1); found {
		t.Error("Search should return false on empty table")
	}
	if st.Delete(1) {
		t.Error("Delete should return false on empty table")
	}
}

func TestSwissTableGrowth(t *testing.T) {
	st := NewSwissTable()
	initialCapacity := st.Capacity()

	// Insert items until resize is triggered
	itemCount := int(float64(initialCapacity)*maxLoadFactor) + 1
	for i := 0; i < itemCount; i++ {
		if !st.Insert(i, fmt.Sprintf("value%d", i)) {
			t.Errorf("Insert failed at index %d", i)
		}
	}

	if st.Capacity() <= initialCapacity {
		t.Error("Table should have grown")
	}

	// Verify all inserted items are still accessible
	for i := 0; i < itemCount; i++ {
		if val, found := st.Search(i); !found || val != fmt.Sprintf("value%d", i) {
			t.Errorf("Item %d not found after resize", i)
		}
	}
}

func TestSwissTableDeletedSlots(t *testing.T) {
	st := NewSwissTable()

	// Insert and delete to create deleted slots
	st.Insert(1, "one")
	st.Insert(2, "two")
	st.Delete(1)

	// Should be able to insert into deleted slot
	if !st.Insert(3, "three") {
		t.Error("Insert into deleted slot failed")
	}
	if val, found := st.Search(3); !found || val != "three" {
		t.Error("Search failed for item in deleted slot")
	}
}
