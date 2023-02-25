// Perennial Systems
package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
Input : A string with a fixed length. Max(1000000)
Output : Top five words which occurred more than once in the string
Ex : "I am a golang developer. Almost all developer like to work in golang. There is a bright future in Golang"
Output :
	golang :3
	a : 2
	developer : 2
*/

/*
It has to case-insensitive
Limit is of 5 words (can be less)
NO punctuations
*/

func main() {
	input := "I am a a a golang is developer. Almost is to to all developer like to to work in is golang. There is in in a bright future in Golang golang"
	wdSlc := strings.Split(strings.ToLower(input), " ")

	wdMp := map[string]int{}

	for _, v := range wdSlc {
		if v[len(v)-1] == '.' {
			v = v[:len(v)-1]
		}
		if _, ok := wdMp[v]; !ok {
			wdMp[v] = 1
		} else {
			wdMp[v] += 1
		}
	}

	cntSet := []int{}

	cntWd := map[int][]string{}
	for k, v := range wdMp {
		if v == 1 {
			continue
		}
		if _, ok := cntWd[v]; !ok {
			cntWd[v] = []string{k}
			cntSet = append(cntSet, v)
		} else {
			cntWd[v] = append(cntWd[v], k)
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(cntSet)))

	count := 0
OuterLoop:
	for _, f := range cntSet {
		for _, v := range cntWd[f] {
			fmt.Println(v, ":", f)
			count++
			if count == 5 {
				break OuterLoop
			}
		}
	}
}
