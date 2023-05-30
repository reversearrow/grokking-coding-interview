package squares_of_sorted_array

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "",
			nums: []int{-14, -1, 0, 2, 3, 10},
			want: []int{0, 1, 4, 9, 100, 196},
		},
		{
			name: "",
			nums: []int{-5, -3, -2, -1},
			want: []int{1, 4, 9, 25},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
