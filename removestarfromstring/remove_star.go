package wigglesort

import (
	"fmt"
)

func RemoveStarFromString(s string) string {
	r := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) != "*" {
			r += string(s[i])
		} else {
			r = r[:len(r)-1]
		}
	}
	return r
}

// Time Complexity :n
// Space: lenght of string

func main() {
	fmt.Println("not doing anything")
}
