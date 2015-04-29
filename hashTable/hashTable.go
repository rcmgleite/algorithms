package hashTable

import "fmt"

////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
/////////////////////// Simple HashTable implementation ////////////////////////
////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////
// Collision resolution strategy: Linear probing
// C++ and Java implementation: http://www.algolist.net/Data_structures/Hash_table/Simple_example

const (
	// TableSize ...
	TableSize = 128
)

// GetHash = hash function
func GetHash(key int) int {
	hash := key % TableSize
	return hash
}

// HashEntry ...
type HashEntry struct {
	key   int // key on hashEntry is used on collision resolution
	value int
}

// GetKey ...
func (he *HashEntry) GetKey() int {
	return he.key
}

// GetValue ...
func (he *HashEntry) GetValue() int {
	return he.value
}

// HashTable ...
type HashTable struct {
	table [TableSize]*HashEntry // array of pointers to HashEntry
}

// Get ...
func (ht *HashTable) Get(key int) int {
	hash := GetHash(key)
	for ht.table[hash] != nil && ht.table[hash].GetKey() != key {
		hash = GetHash(hash + 1)
	}

	if ht.table[hash] == nil {
		return -1
	}
	return ht.table[hash].GetValue()
}

// Put ...
func (ht *HashTable) Put(key, value int) {
	hash := GetHash(key)
	for ht.table[hash] != nil && ht.table[hash].GetKey() != key {
		hash = GetHash(hash + 1)
	}
	ht.table[hash] = &HashEntry{key, value}
}

// PrintAllEntries ... just for debug
func (ht *HashTable) PrintAllEntries() {
	for i := 0; i < TableSize; i++ {
		if ht.table[i] != nil {
			fmt.Printf("ht.table[%d] = %d\n", i, ht.table[i].GetValue())
		}
	}
}
