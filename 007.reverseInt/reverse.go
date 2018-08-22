package main

import (
	"math"
)

//给定一个 32 位有符号整数，将整数中的数字进行反转。
func reverse(x int) int {
	reverseNum := 0
	isNegative := x < 0
	if isNegative {
		x = -x
	}

	for x > 0 {
		reverseNum = reverseNum * 10 + x % 10
		x /= 10
	}

	if isNegative {
		reverseNum = -reverseNum
	}

	if reverseNum > math.MaxInt32 || reverseNum < math.MinInt32 {
		return 0
	}

	return reverseNum
}
