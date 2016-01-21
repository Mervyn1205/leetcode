package main

import "fmt"

func main() {
	var digits = []int{9, 9, 9, 9}
	newDigits := plusOne(digits)
	fmt.Println(newDigits)

	var digits2 = []int{9, 9, 9, 8}
	newDigits = plusOne(digits2)
	fmt.Println(newDigits)

	var digits3 = []int{9, 9, 9, 1}
	newDigits = plusOne(digits3)
	fmt.Println(newDigits)
}

/**
 * Given a non-negative number represented as an array of digits, plus one to the number.
 * The digits are stored such that the most significant digit is at the head of the list.
 */
func plusOne(digits []int) []int {
	sum := 0
	carries := 1
	length := len(digits)
	for i := length - 1; i >= 0 && carries > 0; i-- {
		sum = digits[i] + carries
		digits[i] = sum % 10
		carries = sum / 10
	}

	if carries == 0 {
		return digits
	}

	rst := make([]int, length+1)
	rst[0] = 1
	for i := 1; i < length; i++ {
		rst[i] = digits[i-1]
	}

	return rst
}
