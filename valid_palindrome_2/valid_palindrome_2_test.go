package valid_palindrome_2

import (
	"testing"
)

func Test_validPalindrome(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "madame",
			s:    "madame",
			want: false,
		},
		{
			name: "dead",
			s:    "dead",
			want: true,
		},
		{
			name: "tebbem",
			s:    "tebbem",
			want: false,
		},
		{
			name: "abc",
			s:    "abc",
			want: false,
		},
		{
			name: "madame",
			s:    "madame",
			want: true,
		},
		{
			name: "deeee",
			s:    "deeee",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
