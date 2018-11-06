package p013

import "testing"

func TestRomanToInt(t *testing.T) {
	if romanToInt("III") != 3 {
		t.Error("121 is palindrome")
	}

	if 	romanToInt("IV") != 4 {
		t.Error("0 is palindrome")
	}

	if 	romanToInt("IX") != 9 {
		t.Error("0 is palindrome")
	}

	if 	romanToInt("LVIII") != 58 {
		t.Error("0 is palindrome")
	}

	if 	romanToInt("MCMXCIV") != 1994 {
		t.Error("0 is palindrome")
	}

}
