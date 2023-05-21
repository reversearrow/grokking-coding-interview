package reverse_words_in_a_string

import (
	regexp2 "regexp"
	"strings"
	"unicode"
)

// reverse will reverse any string
func reverseString(s string) string {
	r := []rune(s)
	left := 0
	right := len(r) - 1
	for left < right {
		temp := r[right]
		r[right] = r[left]
		r[left] = temp

		left++
		right--
	}

	return strings.TrimSpace(string(r))
}

func updateStringAtIndex(s string, start int, end int, replacement string) string {
	// Convert the string to a byte slice
	b := []byte(s)

	// Check if the start and end indexes are within bounds
	if start < 0 || end >= len(b) {
		return s // Return the original string if the indexes are invalid
	}

	// Update the slice with the replacement string
	copy(b[start:end+1], []byte(replacement))

	return string(b)
}

func reverseWords(sentence string) string {
	sentence = reverseString(trimSpace(sentence))
	start := 0
	end := 0

	for i, char := range sentence {
		if !unicode.IsSpace(char) {

			continue
		}
		end = i
		sentence = updateStringAtIndex(sentence, start, end, reverseString(sentence[start:end]))
		start = end + 1
	}

	sentence = updateStringAtIndex(sentence, start, len(sentence)-1, reverseString(sentence[start:]))
	return sentence
}

func trimSpace(s string) string {
	s = strings.TrimSpace(s)

	regexp := regexp2.MustCompile(`\s+`)
	return regexp.ReplaceAllString(s, " ")
}
