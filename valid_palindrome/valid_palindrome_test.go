package valid_palindrome

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name        string
		inputString string
		want        bool
	}{
		{
			name:        "kayak",
			inputString: "kayak",
			want:        true,
		},
		{
			name:        "DCBAABCD",
			inputString: "DCBAABCD",
			want:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.inputString); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
