package main

import "testing"

func TestReverse(t *testing.T) {
	result := reverse(1534236469)
	if result != 0 {
		t.Error(result)
		t.Error("result is error, Expect 0")
	}

	result1 := reverse(123)
	if result1 != 321 {
		t.Error(result1)
		t.Error("result is error, Expect 0")
	}

	result2 := reverse(-124)
	if result2 != -421 {
		t.Error(result2)
		t.Error("result is error, Expect 0")
	}

}
