package main

import "fmt"

func main() {
	sli := []int{1, 2, 2, 3, 4}
	length := removeDuplicates(sli)
	fmt.Println(length)

	sli2 := []int{1, 2, 2, 2, 3, 3, 4, 4, 4}
	length2 := removeDuplicates2(sli2, 2)
	fmt.Println(length2)
}

/**
 * Given a sorted array, remove the duplicates in place such that > each element appear only once and return the new length.
 * Do not allocate extra space for another array, you must do this in place with constant memory.
 * For example, Given input array A = [1,1,2],
 * Your function should return length = 2, and A is now [1,2].
 */
func removeDuplicates(srcSlice []int) int {
	j := 0
	length := len(srcSlice)
	for i := 1; i < length; i++ {
		if srcSlice[j] != srcSlice[i] {
			j++
			srcSlice[j] = srcSlice[i]
		}
	}

	srcSlice = srcSlice[:j+1]
	fmt.Println(srcSlice)
	return j + 1
}

/**
 * Follow up for "Remove Duplicates": What if duplicates are allowed at most twice?
 * For example, Given sorted array A = [1,1,1,2,2,3],
 * Your function should return length = 5, and A is now [1,1,2,2,3].
 */
func removeDuplicates2(srcSlice []int, n int) int {
	if n <= 0 {
		return 0
	}

	j := 0
	num := 0
	length := len(srcSlice)
	for i := 1; i < length; i++ {
		if srcSlice[j] != srcSlice[i] {
			j++
			num = 0
			srcSlice[j] = srcSlice[i]
		} else {
			num++
			if num < n {
				j++
				srcSlice[j] = srcSlice[i]
			}
		}
	}

	srcSlice = srcSlice[:j+1]
	fmt.Println(srcSlice)
	return j + 1
}
