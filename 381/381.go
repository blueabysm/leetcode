package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type RandomizedCollection struct {
	datamap  map[string]int
	occurmap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		datamap:  make(map[string]int),
		occurmap: make(map[int]int),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	var result bool
	result = !this.HasValue(val)

	if _, ok := this.occurmap[val]; ok {
		this.occurmap[val]++
	} else {
		this.occurmap[val] = 1
	}

	key := this.hash_code(val)
	this.datamap[key] = val

	return result
}

func (this *RandomizedCollection) HasValue(val int) bool {
	if val, ok := this.occurmap[val]; ok && val > 0 {
		return true
	}
	return false
}

func (this *RandomizedCollection) hash_code(val int) string {
	var occur_count int
	if count, ok := this.occurmap[val]; ok {
		occur_count = count
	} else {
		occur_count = 0
	}

	bytes := []byte(strconv.Itoa(val) + "." + strconv.Itoa(occur_count))
	bytes_key := md5.Sum(bytes)
	return hex.EncodeToString(bytes_key[:])
}

func (this *RandomizedCollection) Get(key string) (int, bool) {
	if val, ok := this.datamap[key]; ok {
		return val, true
	}
	return 0, false
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	key := this.hash_code(val)
	if _, ok := this.datamap[key]; ok {
		delete(this.datamap, key)
		this.occurmap[val]--
		if this.occurmap[val] <= 0 {
			delete(this.occurmap, val)
		}
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedCollection) GetRandom() int {
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

func main() {
	obj := Constructor()
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(1))
	fmt.Println(obj.Insert(2))
	fmt.Println(obj.GetRandom())
	fmt.Println(obj.Remove(1))
	fmt.Println(obj.GetRandom())
}
