package wigglesort

import (
	"fmt"
	"log"
	"testing"

	"golang.org/x/exp/slices"
)

func TestWiggleSort(t *testing.T) {
	arr := []int{3, 5, 1, 10, 30, 20}
	res := WiggleSort(arr)
	fmt.Println(res)
	expected := []int{1, 30, 3, 20, 5, 10}
	equal := slices.Equal(res, expected)
	if !equal {
		log.Fatal("not matching", expected)
	}
}
