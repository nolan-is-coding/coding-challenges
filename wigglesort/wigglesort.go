package wigglesort

import (
	"fmt"
	"sort"
)

// from list of integer
// create a wave sort  up down up down up ...
// sorted
//
//	take the first and then take last elem keep until array is empty.
func WiggleSort(arr []int) []int {
	sort.Ints(arr)
	fmt.Println(arr)
	res := []int{}
	for len(arr) > 0 {
		if len(arr) > 0 {
			res = append(res, arr[0])
			arr = arr[1:]
		}
		if len(arr) > 0 {
			res = append(res, arr[len(arr)-1])
			arr = arr[:len(arr)-1]
		}
	}
	return res
}

// Time Complexity : n log n
// Space: lenght of array

func main() {
	fmt.Println("not doing anything")
}
