// Ingenuity Gaming
package main

import (
	"fmt"
	"sort"
	"math/rand"
	"time"
)

func sortWithMaps(arr [10]int){
    intSlc := make([]int, len(arr))
	intMap := make(map[int]struct{})
	
	for _, val := range(arr){
	    if _, ok := intMap[val]; !ok{
	        intSlc = append(intSlc, val)
	        intMap[val] = struct{}{}
	    }
	}
	sort.Ints(intSlc)
	fmt.Println ("Second largest:", intSlc[len(intSlc)-2])
}

func sortWithArr (arr [10]int){
    sort.Ints(arr[:])
	
	var secLar int
	for secLar = len(arr)-2; arr[secLar] == arr[len(arr)-1]; secLar-- {}
	
	fmt.Println ("Second largest:", arr[secLar])
}

func generate5RandUniqueNum(){
    uniqNumMap := map[int]bool{}
    
    rand.Seed(int64(time.Now().Nanosecond()))
    
    for ;len(uniqNumMap) < 5; {
        num := 7 + rand.Intn(15-7)
        if !uniqNumMap[num]{
            uniqNumMap[num] = true
        }
    }
    
    fmt.Println (uniqNumMap)
    
}

func main() {
	arr := [...]int{2, 2, 4, 5, 3, 7, 9, 8, 8, 9}
	
	if len(arr) < 2{
		fmt.Println (arr[0])
		return
	}
	
	sortWithArr(arr)
	
	sortWithMaps(arr)
    
    generate5RandUniqueNum()
}
