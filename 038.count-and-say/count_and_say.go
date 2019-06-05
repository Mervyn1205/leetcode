package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		str = say(str)
	}
	return str

}

func say(str string) string {
	length := len(str)
	res := ""
	i := 0
	for i < length {
		count := 1
		for (i < length-1) && (str[i] == str[i+1]) {
			count++
			i++
		}
		res += strconv.Itoa(count) + string(str[i])
		i++

	}

	return res
}

func main() {
	fmt.Println(countAndSay(4))
}
