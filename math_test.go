package sugar

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	if got := Max(1, 2); got != 2 {
		t.Errorf("Max(1, 2)=%v, want=%v", got, 2)
	}
	if got := Max(1, math.Inf(1)); !math.IsInf(got, 1) {
		t.Errorf("Max(1, +Inf)=%v, want=%v", got, math.Inf(1))
	}
	if got := Max(1, math.Inf(-1)); got != 1 {
		t.Errorf("Max(1, -Inf)=%v, want=%v", got, 1)
	}
	if got := Max(1, math.NaN()); !math.IsNaN(got) {
		t.Errorf("Max(1, NaN)=%v, want=%v", got, math.NaN())
	}

}
