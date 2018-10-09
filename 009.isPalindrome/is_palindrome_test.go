package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	result := isPalindrome(121)
	if !result {
		t.Error("121 is palindrome")
	}

	result1 := isPalindrome(0)
	if !result1 {
		t.Error("0 is palindrome")
	}

	result2 := isPalindrome(-12)
	if result2 {
		t.Error("-12 is not palindrome")
	}

	result3 := isPalindrome(130)
	if result3 {
		t.Error("130 is not palindrome")
	}

	result4 := isPalindrome(123)
	if result4 {
		t.Error("123 is not palindrome")
	}

}
