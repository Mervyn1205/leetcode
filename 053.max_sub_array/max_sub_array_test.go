package main

import (
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	if  maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) != 6{
		t.Error("result must be 6")
	}

}
