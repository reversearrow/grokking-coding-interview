package valid_palindrome_recursive

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) == 1 {
		return true
	}

	return s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1])
}
