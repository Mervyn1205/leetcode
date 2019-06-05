package main

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	if  countAndSay(1) != "1" {
		t.Error("string must be '1'")
	}

	if  countAndSay(4) != "1211" {
		t.Error("string must be '1211'")
	}

}
