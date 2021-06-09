package main

type MyHashMap struct {
	arr [1000001]int
}

/** Initialize your data structure here. */
func Initialize() MyHashMap {
	arr := [1000001]int{}
	for i := 0; i < len(arr); i++ {
		arr[i] = -1
	}
	return MyHashMap{arr}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	this.arr[key] = value
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	return this.arr[key]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	this.arr[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
