package main

//
//import "math/rand"
//
//func main() {
//	obj := Constructor()
//	println(obj.Insert(1))
//	println(obj.Insert(1))
//	println(obj.Remove(1))
//	println(obj.Remove(1))
//
//	//param_2 :=
//	//param_3 := obj.GetRandom()
//}
//
//type RandomizedSet struct {
//	myMap map[int]int
//}
//
//func Constructor() RandomizedSet {
//	return RandomizedSet{
//		myMap: make(map[int]int),
//	}
//}
//
//func (this *RandomizedSet) Insert(val int) bool {
//	if _, ok := this.myMap[val]; ok {
//		return false
//	}
//	this.myMap[val] = val
//	return true
//}
//
//func (this *RandomizedSet) Remove(val int) bool {
//	if _, ok := this.myMap[val]; !ok {
//		return false
//	}
//	delete(this.myMap, val)
//	return true
//}
//
//func (this *RandomizedSet) GetRandom() int {
//	k := rand.Intn(len(this.myMap))
//	for _, v := range this.myMap {
//		if k == 0 {
//			return v
//		}
//		k--
//	}
//	return 0
//}
//
///**
// * Your RandomizedSet object will be instantiated and called as such:
// * obj := Constructor();
// * param_1 := obj.Insert(val);
// * param_2 := obj.Remove(val);
// * param_3 := obj.GetRandom();
// */
