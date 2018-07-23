package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	len1 := lengthOfLongestSubstring("abcabcbb")
	if len1 != 3 {
		t.Error("result1 is error")
	}

	len2 := lengthOfLongestSubstring("bbbbb")
	if len2 != 1 {
		t.Error("result2 is error")
	}

	len3 := lengthOfLongestSubstring("pwwkew")
	if len3 != 3 {
		t.Error("result3 is error")
	}

}
