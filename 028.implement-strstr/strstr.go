package main

//可以基于KMP字符串匹配算法来处理
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	haystackLen := len(haystack)
	needleLen := len(needle)

	i := 0
	j := 0

	for i < haystackLen && j < needleLen {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			i = i - j + 1 //
			j = 0
		}
	}

	if j == needleLen {
		return i - j
	}

	return -1
}
