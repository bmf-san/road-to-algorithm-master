package hash_table

// ProbingType represents different probing strategies
type ProbingType int

const (
	LinearProbing ProbingType = iota
	QuadraticProbing
	DoubleHashing
)

// Item represents a key-value pair in the hash table
type Item struct {
	Key   int
	Value interface{}
	IsSet bool // flag to track if the slot is used
}

// HashTable represents the open addressing hash table
type HashTable struct {
	items       []Item
	size        int
	capacity    int
	probingType ProbingType
}

// NewHashTable creates a new hash table with given capacity and probing strategy
func NewHashTable(capacity int, probingType ProbingType) *HashTable {
	return &HashTable{
		items:       make([]Item, capacity),
		size:        0,
		capacity:    capacity,
		probingType: probingType,
	}
}

// hash calculates the primary hash value for a key
func (ht *HashTable) hash(key int) int {
	return key % ht.capacity
}

// secondaryHash calculates a secondary hash value for double hashing
func (ht *HashTable) secondaryHash(key int) int {
	// Common secondary hash function: h2(key) = PRIME - (key % PRIME)
	// Where PRIME is smaller than table size
	prime := ht.capacity - 2 // Using a prime number smaller than table size
	return prime - (key % prime)
}

// probe calculates the next probe position based on the selected probing strategy
func (ht *HashTable) probe(pos, attempt int, key int) int {
	switch ht.probingType {
	case LinearProbing:
		return (pos + attempt) % ht.capacity
	case QuadraticProbing:
		// Using quadratic formula: h(k,i) = (h(k) + c1*i + c2*i^2) % m
		// Here using c1=c2=1 for simplicity
		return (pos + attempt + attempt*attempt) % ht.capacity
	case DoubleHashing:
		step := ht.secondaryHash(key)
		return (pos + attempt*step) % ht.capacity
	default:
		return (pos + attempt) % ht.capacity
	}
}

// Insert adds a new key-value pair to the hash table
func (ht *HashTable) Insert(key int, value interface{}) bool {
	if ht.size >= ht.capacity {
		return false // table is full
	}

	pos := ht.hash(key)
	attempt := 0

	for attempt < ht.capacity {
		currentPos := ht.probe(pos, attempt, key)
		if !ht.items[currentPos].IsSet {
			// Found empty slot
			ht.items[currentPos] = Item{
				Key:   key,
				Value: value,
				IsSet: true,
			}
			ht.size++
			return true
		}
		if ht.items[currentPos].Key == key {
			// Update existing key
			ht.items[currentPos] = Item{
				Key:   key,
				Value: value,
				IsSet: true,
			}
			return true
		}
		attempt++
	}

	return false // table is full after probing
}

// Search looks up a value by key in the hash table
func (ht *HashTable) Search(key int) (interface{}, bool) {
	pos := ht.hash(key)
	attempt := 0

	for attempt < ht.capacity {
		currentPos := ht.probe(pos, attempt, key)
		if !ht.items[currentPos].IsSet {
			return nil, false // key not found
		}
		if ht.items[currentPos].Key == key {
			return ht.items[currentPos].Value, true
		}
		attempt++
	}

	return nil, false // key not found after probing
}

// Delete removes a key-value pair from the hash table
func (ht *HashTable) Delete(key int) bool {
	pos := ht.hash(key)
	attempt := 0

	for attempt < ht.capacity {
		currentPos := ht.probe(pos, attempt, key)
		if !ht.items[currentPos].IsSet {
			return false // key not found
		}
		if ht.items[currentPos].Key == key {
			ht.items[currentPos] = Item{} // Clear the slot
			ht.size--
			return true
		}
		attempt++
	}

	return false // key not found after probing
}

// Size returns the number of items in the hash table
func (ht *HashTable) Size() int {
	return ht.size
}

// Capacity returns the total capacity of the hash table
func (ht *HashTable) Capacity() int {
	return ht.capacity
}
