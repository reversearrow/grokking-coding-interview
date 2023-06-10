package pair_with_three_sum

import (
	"testing"
)

func Test_findSumOfThree(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "",
			nums:   []int{3, 7, 1, 2, 8, 4, 5},
			target: 10,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: 0,
			want:   true,
		},
		{
			name:   "",
			nums:   []int{-1, 2, 1, -4, 5, -3},
			target: -8,
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSumOfThree(tt.nums, tt.target); got != tt.want {
				t.Errorf("findSumOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
