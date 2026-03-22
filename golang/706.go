//go:build ignore

package main

// Design a HashMap without using any built-in hash table libraries.
//
// Implement the MyHashMap class:
//
// MyHashMap() initializes the object with an empty map.
// void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
// int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
// void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.
//
//
// Example 1:
//
// Input
// ["MyHashMap", "put", "put", "get", "get", "put", "get", "remove", "get"]
// [[], [1, 1], [2, 2], [1], [3], [2, 1], [2], [2], [2]]
// Output
// [null, null, null, 1, -1, null, 1, null, -1]
//
// Explanation
// MyHashMap myHashMap = new MyHashMap();
// myHashMap.put(1, 1); // The map is now [[1,1]]
// myHashMap.put(2, 2); // The map is now [[1,1], [2,2]]
// myHashMap.get(1);    // return 1, The map is now [[1,1], [2,2]]
// myHashMap.get(3);    // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
// myHashMap.put(2, 1); // The map is now [[1,1], [2,1]] (i.e., update the existing value)
// myHashMap.get(2);    // return 1, The map is now [[1,1], [2,1]]
// myHashMap.remove(2); // remove the mapping for 2, The map is now [[1,1]]
// myHashMap.get(2);    // return -1 (i.e., not found), The map is now [[1,1]]
//
//
// Constraints:
//
// 0 <= key, value <= 106
// At most 104 calls will be made to put, get, and remove.

type Node struct {
	Val  int
	Next *Node
	Key  int
}

type MyHashMap struct {
	HashMap map[int]*Node
}

func Constructor() MyHashMap {
	// why 1000 because
	// key and value only up to 10^6, but 10^4 operation can made.
	// if we choose 10^2 or 10^1 -> which make collision for linked list will be a lot
	// if we choose 10^4 -> which make hash map capacity big, and make extra space for hash map which is not necessary.
	// 10^3 is perfect for this scenario, the linked list is work well

	m := make(map[int]*Node, 1000)
	for i := range 1000 {
		m[i] = &Node{
			Val: -1,
			Key: -1,
		}
	}

	return MyHashMap{
		HashMap: m,
	}
}

func (this *MyHashMap) Hash(key int) int {
	return key % len(this.HashMap)
}

func (this *MyHashMap) Put(key int, value int) {
	idx := this.Hash(key)

	curr := this.HashMap[idx]
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next.Val = value
			return
		}
		curr = curr.Next
	}
	curr.Next = &Node{
		Key: key, Val: value,
	}
}

func (this *MyHashMap) Get(key int) int {
	idx := this.Hash(key)

	curr := this.HashMap[idx].Next
	for curr != nil {
		if curr.Key == key {
			return curr.Val
		}
		curr = curr.Next
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	idx := this.Hash(key)

	curr := this.HashMap[idx]
	for curr.Next != nil {
		if curr.Next.Key == key {
			curr.Next = curr.Next.Next
			return
		}
		curr = curr.Next
	}
}

func main() {
	/**
	 * Your MyHashMap object will be instantiated and called as such:
	 * obj := Constructor();
	 * obj.Put(key,value);
	 * param_2 := obj.Get(key);
	 * obj.Remove(key);
	 */
}
