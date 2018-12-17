package main

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	if  removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}) != 5 {
		t.Error("length should be 5")
	}

	if  removeDuplicates([]int{1,1,2}) != 2 {
		t.Error("length should be 5")
	}

}
