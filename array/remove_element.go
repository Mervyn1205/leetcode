/**
 * http://www.lintcode.com/zh-cn/problem/remove-element/
 * 给定一个数组和一个值，在原地删除与值相同的数字，返回新数组的长度。
 * 元素的顺序可以改变，并且对新的数组不会有影响。
 */
package array

/**
 * eg:
 *  a := []int{1, 2, 2, 3, 2, 4}
 *	length := array.RemoveElement(a, 2)
 *	fmt.Println(length)
 */
func RemoveElement(srcSlice []int, elem int) int {
	length := len(srcSlice)
	j := 0
	for i := 0; i < length; i++ {
		if srcSlice[i] == elem {
			continue
		}

		srcSlice[j] = srcSlice[i]
		j++
	}

	return j
}
