package removestarfromstring

import (
	"fmt"
	"log"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	ts := "this*is**astring*"
	res := RemoveStarFromString(ts)
	fmt.Println(res)
	expected := "thiastrin"
	if res != expected {
		log.Fatal("not matching", expected)
	}
}
