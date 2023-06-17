package sum_numbers_in_an_array

import "testing"

func Test_sum(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{
			name: "",
			arr:  []int{1, 3},
			want: 4,
		},
		{
			name: "",
			arr:  []int{1, 3, 5, 9},
			want: 18,
		},
		{
			name: "",
			arr:  []int{1, 10, 5, 9},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.arr); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
