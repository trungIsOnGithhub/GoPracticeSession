package hash_map

import (
	"fmt"
)

const ArraySize = 8

// Hash table has an array internally
type HashTable struct {
	array [ArraySize]*bucket
}

// Each position in array is linked list bucket
type bucket struct {
	llHead *bucketNode
}

// LList Bucket has bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

// HashTable construtor
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}
func (h *HashTable) Delete(key string)  {
	index := hash(key)
	h.array[index].delete(key)
}

// insert for bucket at each position
func (b *bucket) insert(k string) {
	if b.search(k) { return }

	newNode := &bucketNode{ key: string }
	newNode.next = b.head
	b.head = newNode
}

// search for a key in bucket
func (b* bucket) search(key string) bool {
	currentNode := b.head
	
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	
	return false
}

// delete a key in bucket
func (b* bucket) delete(key string) {
	if b.head == nil { return }
	
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	
	currentNode := b.head
	
	for currentNode.next != nil {
		if currentNode.next.key == k {
			currentNode.next = currentNode.next.next
		}
		currentNode = currentNode.next
	}
}

// hash function
func hash(key string) int {
	sum := 0
	for _, c := range key {
		sum += int(c)
	}
	return sum % ArraySize
}

func main() {
	testHashTab := Init()
	testHashTab.Insert("Duy")
	testHashTab.Insert("Hieu")
	testHashTab.Insert("Tuan")
	testHashTab.Insert("Trung")
	testHashTab.Insert("Vu")
	testHashTab.Delete("Tuan")
	testHashTab.Delete("Duy")
	testHashTab.Search("Trung")
	testHashTab.Search("Hieu")
}