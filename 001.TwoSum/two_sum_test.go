package main

import "testing"

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	result := twoSum1(nums, target)

	if result[0] != 0 || result[1] != 2 {
		t.Error(result)
		t.Error("result is error:")
	}

}

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	result := twoSum1(nums, target)

	if result[0] != 0 || result[1] != 2 {
		t.Error(result)
		t.Error("result is error:")
	}

}

func TestTwoSum3(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 13
	result := twoSum1(nums, target)

	if result[0] != 0 || result[1] != 2 {
		t.Error(result)
		t.Error("result is error:")
	}

}
