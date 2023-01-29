package main

func main() {

}

type MyHashMap struct {
	HashMap map[int]int
}

func Constructor706() MyHashMap {
	return MyHashMap{HashMap: make(map[int]int)}
}

func (this *MyHashMap) Put(key int, value int) {
	this.HashMap[key] = value
}

func (this *MyHashMap) Get(key int) int {
	if val, ok := this.HashMap[key]; ok {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	delete(this.HashMap, key)
}
