package wigglesort

import (
	"fmt"
	"sort"
)

func WiggleSort(arr []int) []int {
	sort.Ints(arr)
	fmt.Println("arr", arr)
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
	fmt.Println(res)
	return res
}

func main() {
	fmt.Println("not doing anything")
}
