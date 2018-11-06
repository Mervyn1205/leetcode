package main

func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return result
	}

	minStrKey := 0
	minLen := len(strs[0])
	for key, str := range strs {
		strLen := len(str)

		if strLen < minLen {
			minLen = strLen
			minStrKey = key
		}
	}

	isSame := true
	for key, val := range strs[minStrKey] {
		for k, str := range strs {
			if k == minStrKey {
				continue
			}
			if rune(str[key]) != val {
				isSame = false
				break
			}
		}
		if !isSame {
			break
		}
		result += string(val)
	}

	return result
}
