package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	if  !isValid("()") {
		t.Error("() is valid")
	}

	if  !isValid("()[]{}") {
		t.Error("() is valid")
	}

	if  isValid("(]") {
		t.Error("(] is not valid")
	}
	if  isValid("([)]") {
		t.Error("([)] is not valid")
	}

	if  !isValid("{[]}") {
		t.Error("{[]} is valid")
	}


}
