package pair_with_target_sum

import (
	"reflect"
	"testing"
)

func Test_pairWithTargetSum(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pairWithTargetSum(tt.input, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pairWithTargetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
