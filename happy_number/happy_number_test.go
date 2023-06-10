package happy_number

import (
	"testing"
)

func Test_isHappy(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{
			name: "",
			num:  18,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.num); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
