package valid_palindrome

func isPalindrome(inputString string) bool {
	start := 0
	end := len(inputString) - 1

	for {
		if start >= len(inputString)/2 && end <= len(inputString)/2 {
			return true
		}

		if inputString[start] != inputString[end] {
			return false
		}

		start++
		end--
	}

	return true
}
