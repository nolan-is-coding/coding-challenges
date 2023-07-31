package main

import (
	"fmt"
	"sort"
)

// find the min and max sum of 4 elements in an slice
// e.g [1,2,3,4,5] => 10 , 14

func miniMaxSum(arr []int32) (int64, int64) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	var minSum int64
	var maxSum int64
	for i := 0; i < len(arr) && i < 4; i++ {
		minSum += int64(arr[i])
	}
	for i := len(arr) - 1; len(arr)-1 > 0 && i != len(arr)-5; i-- {
		maxSum += int64(arr[i])
	}
	return minSum, maxSum
}

func main() {
	fmt.Println("doing nothing")
}
