package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	if  longestCommonPrefix([]string{"flower","flow","flight"}) != "fl" {
		t.Error("error")
	}

	if  longestCommonPrefix([]string{"dog","racecar","car"}) != "" {
		t.Error("error 2")
	}

}
