package numbers

import (
	"reflect"
	"testing"
)

func TestMax(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		// TODO: Add test cases.
		{"test max all zeros", args{0, 0}, 0},
		{"test max positive", args{1, 2}, 2},
		{"test max negtive", args{-1, -2}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
