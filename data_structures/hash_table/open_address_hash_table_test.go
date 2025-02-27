package hash_table

import "testing"

func testHashTableOperations(t *testing.T, probingType ProbingType, name string) {
	// Create a new hash table
	ht := NewHashTable(5, probingType)

	// Test initial state
	if ht.Size() != 0 {
		t.Errorf("[%s] Expected size 0, got %d", name, ht.Size())
	}
	if ht.Capacity() != 5 {
		t.Errorf("[%s] Expected capacity 5, got %d", name, ht.Capacity())
	}

	// Test Insert
	if !ht.Insert(1, "one") {
		t.Errorf("[%s] Insert failed", name)
	}
	if !ht.Insert(2, "two") {
		t.Errorf("[%s] Insert failed", name)
	}
	if ht.Size() != 2 {
		t.Errorf("[%s] Expected size 2, got %d", name, ht.Size())
	}

	// Test Search
	if val, found := ht.Search(1); !found || val != "one" {
		t.Errorf("[%s] Search failed for key 1", name)
	}
	if val, found := ht.Search(2); !found || val != "two" {
		t.Errorf("[%s] Search failed for key 2", name)
	}
	if _, found := ht.Search(3); found {
		t.Errorf("[%s] Search should not find non-existent key", name)
	}

	// Test collision handling
	if !ht.Insert(6, "six") { // 6 % 5 = 1, should handle collision
		t.Errorf("[%s] Insert failed on collision", name)
	}
	if val, found := ht.Search(6); !found || val != "six" {
		t.Errorf("[%s] Search failed for collided key", name)
	}

	// Test Delete
	if !ht.Delete(1) {
		t.Errorf("[%s] Delete failed", name)
	}
	if _, found := ht.Search(1); found {
		t.Errorf("[%s] Search found deleted key", name)
	}
	if ht.Size() != 2 {
		t.Errorf("[%s] Expected size 2 after delete, got %d", name, ht.Size())
	}

	// Test table full condition
	ht.Insert(3, "three")
	ht.Insert(4, "four")
	ht.Insert(5, "five")
	if ht.Insert(7, "seven") { // Should fail as table is full
		t.Errorf("[%s] Insert should fail when table is full", name)
	}
}

func TestLinearProbing(t *testing.T) {
	testHashTableOperations(t, LinearProbing, "Linear Probing")
}

func TestQuadraticProbing(t *testing.T) {
	testHashTableOperations(t, QuadraticProbing, "Quadratic Probing")
}

func TestDoubleHashing(t *testing.T) {
	testHashTableOperations(t, DoubleHashing, "Double Hashing")
}

func TestHashTableUpdates(t *testing.T) {
	// Test with each probing type
	probingTypes := []struct {
		pType ProbingType
		name  string
	}{
		{LinearProbing, "Linear Probing"},
		{QuadraticProbing, "Quadratic Probing"},
		{DoubleHashing, "Double Hashing"},
	}

	for _, pt := range probingTypes {
		t.Run(pt.name, func(t *testing.T) {
			ht := NewHashTable(3, pt.pType)

			// Test updating existing key
			ht.Insert(1, "one")
			ht.Insert(1, "ONE")
			if val, found := ht.Search(1); !found || val != "ONE" {
				t.Errorf("[%s] Update existing key failed", pt.name)
			}
			if ht.Size() != 1 {
				t.Errorf("[%s] Size should be 1 after update, got %d", pt.name, ht.Size())
			}
		})
	}
}

func TestHashTableEmptyOperations(t *testing.T) {
	// Test with each probing type
	probingTypes := []struct {
		pType ProbingType
		name  string
	}{
		{LinearProbing, "Linear Probing"},
		{QuadraticProbing, "Quadratic Probing"},
		{DoubleHashing, "Double Hashing"},
	}

	for _, pt := range probingTypes {
		t.Run(pt.name, func(t *testing.T) {
			ht := NewHashTable(3, pt.pType)

			// Test operations on empty table
			if _, found := ht.Search(1); found {
				t.Errorf("[%s] Search should return false on empty table", pt.name)
			}
			if ht.Delete(1) {
				t.Errorf("[%s] Delete should return false on empty table", pt.name)
			}
		})
	}
}
