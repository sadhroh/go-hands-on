package main

import (
	"fmt"
)

const input = "aabbbbc   cccdddd"

type BestMap struct {
	keys []rune

	charMap map[rune]*meta
}

type meta struct {
	mapVal int
	keyIdx int
}

func (b BestMap) Keys() []rune {
	return b.keys
}

func (b BestMap) Exists(k rune) bool {
	_, ok := b.charMap[k]
	return ok
}

func (b BestMap) Get(k rune) (int, bool) {
	v, ok := b.charMap[k]
	if ok {
		return v.mapVal, ok
	}
	return 0, ok
}

func (b *BestMap) Put(k rune, v int) {
	m := &meta{
		mapVal: v,
	}
	if !b.Exists(k) {
		b.keys = append(b.keys, k)
		m.keyIdx = len(b.keys) - 1
	} else {
		m.keyIdx = b.charMap[k].keyIdx
	}
	b.charMap[k] = m
}

func (b *BestMap) Del(k rune) bool {
	if !b.Exists(k) {
		return false
	}
	m := b.charMap[k]
	b.keys = append(b.keys[:m.keyIdx], b.keys[m.keyIdx+1:]...)
	delete(b.charMap, k)
	return true
}

func (b BestMap) String() string {
	var data string
	for k, v := range b.charMap {
		if data != "" {
			data += "\n"
		}
		data = fmt.Sprintf("%s%s:%d", data, string(k), v.mapVal)
	}
	return data
}

func NewMap() *BestMap {
	return &BestMap{
		keys:    nil,
		charMap: make(map[rune]*meta),
	}
}

func main() {
	charMap := NewMap()

	var max int

	for _, c := range input {
		if c == ' ' {
			continue
		}
		if !charMap.Exists(c) {
			charMap.Put(c, 1)
		} else if v, ok := charMap.Get(c); ok {
			charMap.Put(c, v+1)
		}

		v, _ := charMap.Get(c)
		if max < v {
			max = v
		}
	}

	var (
		count int
		c     rune
	)
	for _, k := range charMap.Keys() {
		v, _ := charMap.Get(k)
		if v != max && count < v {
			count = v
			c = k
		}
	}

	fmt.Println(string(c))
}
