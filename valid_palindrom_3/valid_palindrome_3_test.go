package valid_palindrom_3

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "A man, a plan, a canal: Panama",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "0P",
			s:    "0P",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
