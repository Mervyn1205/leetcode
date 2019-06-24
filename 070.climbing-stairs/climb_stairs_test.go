package main

import (
	"testing"
)

func TestClimbStairs1(t *testing.T) {
	if  climbStairs1(1) != 1 {
		t.Error("must be 1")
	}

	if  climbStairs1(2) != 2 {
		t.Error("must be 1")
	}

	if  climbStairs1(3) != 3 {
		t.Error("must be 1")
	}


}


func TestClimbStairs2(t *testing.T) {
	if  climbStairs2(1) != 1 {
		t.Error("must be 1")
	}

	if  climbStairs2(2) != 2 {
		t.Error("must be 1")
	}

	if  climbStairs2(3) != 3 {
		t.Error("must be 1")
	}


}
