package main

import (
	"fmt"
	"math/rand"
)

func main() {
	set := Constructor()
	for i := 0; i < 10; i++ {
		set.Insert(i)
	}

	for i := 0; i < 9; i++ {
		set.Remove(i)
	}
	fmt.Printf("%v\n", set.GetRandom())
}

type RandomizedSet struct {
	m map[int]int // [val]idx
	a []int
}

func Constructor() RandomizedSet {
	n := 100
	return RandomizedSet{
		m: make(map[int]int, n),
		a: make([]int, 0, n),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.m[val]

	if ok {
		return false
	}

	this.m[val] = len(this.a)
	this.a = append(this.a, val)

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.m[val]

	if !ok {
		return false
	}

	if len(this.a) == 1 {
		delete(this.m, val)
		this.a = this.a[:0]
		fmt.Printf("set %v\n", this)
		return true
	}

	lastIdx := len(this.a) - 1
	if idx != lastIdx {
		// switch last and deleted
		this.a[lastIdx], this.a[idx] = this.a[idx], this.a[lastIdx]
		this.m[this.a[idx]] = idx
	}
	// update val in the map
	delete(this.m, this.a[lastIdx])

	// cut off last item
	this.a = this.a[:lastIdx]
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
