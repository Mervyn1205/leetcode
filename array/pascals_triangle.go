package main

import "fmt"

func main() {
	PascalTriangle(6)
	GetRow(5)
}

/**
 * Given numRows, generate the first numRows of Pascal's triangle.
 *	For example, given numRows = 5, Return
 *
 * [
 *      [1],
 *     [1,1],
 *    [1,2,1],
 *   [1,3,3,1],
 *  [1,4,6,4,1]
 * ]
 */
func PascalTriangle(row int) {
	data := make([][]int, row)
	for i := 0; i < row; i++ {
		data[i] = make([]int, i+1)
		data[i][0] = 1
		data[i][i] = 1
		for j := 1; j < i; j++ {
			data[i][j] = data[i-1][j-1] + data[i-1][j]
		}
	}

	for _, val := range data {
		fmt.Println(val)
	}
}

/**
 * Given an index k, return the kth row of the Pascal's triangle.
 * For example, given k = 3, Return [1,3,3,1].
 */
func GetRow(index int) {
	row := make([]int, index)
	row[0] = 1
	for i := 2; i < index+1; i++ {
		for j := i - 1; j >= 1; j-- {
			row[j] = row[j] + row[j-1]
		}
	}

	fmt.Println(row)
}
