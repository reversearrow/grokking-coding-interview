package valid_palindrome_2

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			if valid(s, left+1, right) {
				return true
			}

			if valid(s, left, right-1) {
				return true
			}

			return false
		}

		left++
		right--
	}

	return true
}

func valid(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}
