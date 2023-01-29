package main

func main() {

}

type MyHashSet struct {
	Hashset map[int]int
}

func Constructor705() MyHashSet {
	return MyHashSet{Hashset: make(map[int]int)}
}

func (this *MyHashSet) Add(key int) {
	this.Hashset[key]++
}

func (this *MyHashSet) Remove(key int) {
	delete(this.Hashset, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.Hashset[key]
	return ok
}
