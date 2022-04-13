package sugar

import "testing"

func TestCondition(t *testing.T) {
	type args struct {
		con bool
		v1  int
		v2  int
	}
	tests := []struct {
		name string
		arg  args
		want int
	}{
		{"1<2?1:2", args{1 < 2, 1, 2}, 1},
		{"1>2?1:2", args{1 > 2, 1, 2}, 2},
		{"true?true:false", args{true, 1, 0}, 1},
		{"a==b?a:b", args{"a" == "b", 'a', 'b'}, 'b'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Condition(tt.arg.con, tt.arg.v1, tt.arg.v2); got != tt.want {
				t.Errorf("got(%v, %v, %v)=%v, want=%v",
					tt.arg.con, tt.arg.v1, tt.arg.v2, got, tt.want)
			}
		})
	}
}
