package findthehighestaltitude

import "testing"

func TestFindTheHighestAltitude(t *testing.T) {
	input := []int{-5, 1, 5, 0, -7}
	r := FindTheHighestAltitude(input)
	if r != 1 {
		t.Fatal("failed")
	}
}
