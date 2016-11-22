package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type RandomizedSet struct {
	datamap map[string]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		datamap: make(map[string]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	key := this.hash_code(val)
	if _, ok := this.Get(key); ok {
		return false
	}
	this.datamap[key] = val
	return true
}

func (this *RandomizedSet) hash_code(val int) string {
	bytes := []byte(strconv.Itoa(val))
	bytes_key := md5.Sum(bytes)
	return hex.EncodeToString(bytes_key[:])
}

func (this *RandomizedSet) Get(key string) (int, bool) {
	if val, ok := this.datamap[key]; ok {
		return val, true
	}
	return 0, false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	key := this.hash_code(val)
	if _, ok := this.datamap[key]; ok {
		delete(this.datamap, key)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.datamap))
	var value int
	for _, v := range this.datamap {
		value = v
		if index <= 0 {
			break
		}
		index--
	}
	return value
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(-1))
	fmt.Println(obj.Remove(-2))
	fmt.Println(obj.Insert(-2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(-1))
	fmt.Println(obj.Insert(-2))
	fmt.Println(obj.GetRandom())
}
