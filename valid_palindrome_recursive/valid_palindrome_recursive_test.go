package valid_palindrome_recursive

import "testing"

func Test_valid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "",
			input: "welcome",
			want:  false,
		},
		{
			name:  "",
			input: "aba",
			want:  true,
		},
		{
			name:  "kayak",
			input: "kayak",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.input); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
