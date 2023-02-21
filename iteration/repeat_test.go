package iteration

import "testing"

func TestRepeat(t *testing.T) {
	type args struct {
		character string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"number repeat", args{"1"}, "11111"},
		{"alpha repeat", args{"a"}, "aaaaa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Repeat(tt.args.character); got != tt.want {
				t.Errorf("Repeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
