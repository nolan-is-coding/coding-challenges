package main

import "testing"

func TestMinMaxSum1(t *testing.T) {
	input := []int32{1, 2, 3, 4, 5}
	a, b := miniMaxSum(input)
	if a != 10 || b != 14 {
		t.Fatal("failed test 1")
	}
}

func TestMinMaxSum2(t *testing.T) {
	input := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
	a, b := miniMaxSum(input)
	if a != 2063136757 || b != 2744467344 {
		t.Fatal("failed test 2")
	}
}
