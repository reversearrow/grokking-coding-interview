package remove_duplicates

import (
	"reflect"
	"testing"
)

func Test_remove(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		input []int
		index int
		want  []int
	}{
		{
			name:  "[2, 3, 3, 3, 6, 9, 9]",
			input: []int{2, 3, 3, 3, 6, 9, 9},
			index: 4,
			want:  []int{2, 3, 3, 3, 9, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := remove(tt.input, tt.index)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "[2, 3, 3, 3, 6, 9, 9]",
			input: []int{2, 2, 3, 3, 3, 6, 9, 9},
			want:  []int{2, 3, 6, 9},
		},
		{
			name:  "[2, 2,  3, 3, 3, 6, 9, 9]",
			input: []int{2, 2, 3, 3, 3, 6, 9, 9},
			want:  []int{2, 3, 6, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
