package reverse_words_in_a_string

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		name     string
		sentence string
		want     string
	}{
		{
			name:     "We love GoLang",
			sentence: "We love GoLang",
			want:     "GoLang love We",
		},
		{
			name:     "Hello     World",
			sentence: "Hello     World",
			want:     "World Hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.sentence); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
