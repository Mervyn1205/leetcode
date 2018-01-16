package main

import "fmt"

func main() {
	nums := []int {2, 7, 11, 15}
	target := 13
	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))

	target = 6
	nums = []int {3, 3}
	fmt.Println(twoSum1(nums, target))
	fmt.Println(twoSum2(nums, target))
	fmt.Println(twoSum3(nums, target))

}

/**
 * Given an array of integers, return indices of the two numbers such that they add up to a specific target.
 * You may assume that each input would have exactly one solution, and you may not use the same element twice.
 * https://leetcode.com/problems/two-sum/description/
 */

func twoSum1(nums []int, target int) []int {
	result := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				result[0] = i
				result[1] = j
			}
		}
	}

	return result
}

func twoSum2(nums []int, target int) []int {
	result := make([]int, 2)
	numsMap := make(map[int]int)

	for k, v := range nums {
		numsMap[v] = k
	}

	for k, v := range nums {
		complement := target - v
		key, ok := numsMap[complement];
		if ok && key != k {
			result[0] = k
			result[1] = key
			break
		}
	}
	return result
}

func twoSum3(nums []int, target int) []int {
	result := make([]int, 2)
	numsMap := make(map[int]int)

	for k, v := range nums {
		complement := target - v
		key, ok := numsMap[complement];
		if ok {
			result[0] = key
			result[1] = k
			break
		}

		numsMap[v] = k
	}
	return result
}


