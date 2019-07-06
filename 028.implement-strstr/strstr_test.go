package main

import (
	"testing"
)

func TestStrstr(t *testing.T) {
	if strStr("mississippi", "issipi") != -1 {
		t.Error("result1 should be -1")
	}

	if strStr("hello", "ll") != 2 {
		t.Error("result2 should be 2")
	}

	if strStr("aaaaa", "bba") != -1 {
		t.Error("result3 should be -1")
	}
}
