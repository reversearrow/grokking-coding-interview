package valid_palindrom_3

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	left, right := 0, len(s)-1

	for left < right {
	}

}

func valid(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
