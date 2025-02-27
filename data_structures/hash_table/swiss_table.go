package hash_table

const (
	groupSize     = 16 // グループサイズ（SSE2を想定）
	maxLoadFactor = 0.875
	growthFactor  = 2
	minCapacity   = groupSize
	metaEmpty     = 0    // 空スロット
	metaDeleted   = 1    // 削除済みスロット
	metaFull      = 0xFF // 使用中スロット
)

// Item はハッシュテーブルの要素を表す
type SwissTableItem struct {
	Key   int
	Value interface{}
}

// SwissTable はスイステーブルの実装
type SwissTable struct {
	metadata []byte           // メタデータビット
	items    []SwissTableItem // 実データ
	size     int              // 使用中の要素数
	capacity int              // テーブルの容量
}

// NewSwissTable は新しいスイステーブルを作成する
func NewSwissTable() *SwissTable {
	capacity := minCapacity
	return &SwissTable{
		metadata: make([]byte, capacity),
		items:    make([]SwissTableItem, capacity),
		size:     0,
		capacity: capacity,
	}
}

// hash は主ハッシュ値を計算する
func (st *SwissTable) hash(key int) uint64 {
	// FNV-1a hash
	hash := uint64(14695981039346656037)
	bytes := []byte{
		byte(key),
		byte(key >> 8),
		byte(key >> 16),
		byte(key >> 24),
	}
	for _, b := range bytes {
		hash ^= uint64(b)
		hash *= 1099511628211
	}
	return hash
}

// probeDistance はプローブ距離を計算する
func (st *SwissTable) probeDistance(hash uint64, pos int) int {
	ideal := int(hash % uint64(st.capacity))
	if pos >= ideal {
		return pos - ideal
	}
	return st.capacity - ideal + pos
}

// findGroup は指定されたキーが所属するグループを探す
func (st *SwissTable) findGroup(key int) (int, bool) {
	if st.size == 0 {
		return 0, false
	}

	hash := st.hash(key)
	pos := int(hash % uint64(st.capacity))
	distance := 0

	for {
		// グループ単位で探索
		for i := 0; i < groupSize && pos+i < st.capacity; i++ {
			currentPos := pos + i
			if st.metadata[currentPos] == metaEmpty {
				return currentPos, false
			}
			if st.metadata[currentPos] == metaFull && st.items[currentPos].Key == key {
				return currentPos, true
			}
		}

		distance += groupSize
		if distance >= st.capacity {
			return 0, false
		}
		pos = (pos + groupSize) % st.capacity
	}
}

// resize はテーブルのサイズを変更する
func (st *SwissTable) resize() {
	oldItems := st.items
	oldMeta := st.metadata
	oldCap := st.capacity

	st.capacity *= growthFactor
	st.metadata = make([]byte, st.capacity)
	st.items = make([]SwissTableItem, st.capacity)
	st.size = 0

	// 既存の要素を再挿入
	for i := 0; i < oldCap; i++ {
		if oldMeta[i] == metaFull {
			st.Insert(oldItems[i].Key, oldItems[i].Value)
		}
	}
}

// Insert は要素を挿入する
func (st *SwissTable) Insert(key int, value interface{}) bool {
	// ロードファクターをチェック
	if float64(st.size+1)/float64(st.capacity) > maxLoadFactor {
		st.resize()
	}

	hash := st.hash(key)
	pos := int(hash % uint64(st.capacity))
	distance := 0

	// 既存のキーを更新
	if existingPos, found := st.findGroup(key); found {
		st.items[existingPos].Value = value
		return true
	}

	// 新しい要素を挿入
	for {
		if st.metadata[pos] == metaEmpty || st.metadata[pos] == metaDeleted {
			st.metadata[pos] = metaFull
			st.items[pos] = SwissTableItem{Key: key, Value: value}
			st.size++
			return true
		}

		// ロビンフッド法: より長くプローブしている要素と交換
		existingDistance := st.probeDistance(st.hash(st.items[pos].Key), pos)
		if distance > existingDistance {
			// 要素を交換
			newMeta := metaFull
			st.metadata[pos] = byte(newMeta)

			newItem := SwissTableItem{Key: key, Value: value}
			oldItem := st.items[pos]
			st.items[pos] = newItem

			key = oldItem.Key
			value = oldItem.Value
			distance = existingDistance
		}

		pos = (pos + 1) % st.capacity
		distance++
		if distance >= st.capacity {
			return false // テーブルが満杯
		}
	}
}

// Search はキーに対応する値を探す
func (st *SwissTable) Search(key int) (interface{}, bool) {
	if pos, found := st.findGroup(key); found {
		return st.items[pos].Value, true
	}
	return nil, false
}

// Delete は要素を削除する
func (st *SwissTable) Delete(key int) bool {
	if pos, found := st.findGroup(key); found {
		st.metadata[pos] = metaDeleted
		st.items[pos] = SwissTableItem{}
		st.size--
		return true
	}
	return false
}

// Size は現在の要素数を返す
func (st *SwissTable) Size() int {
	return st.size
}

// Capacity は現在の容量を返す
func (st *SwissTable) Capacity() int {
	return st.capacity
}
