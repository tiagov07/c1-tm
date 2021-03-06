package tdd

import "testing"

func TestFactorial(t *testing.T) {
	test := []struct {
		arg  int
		want int
	}{{0, 1}, {5, 120}}
	for i, d := range test {
		got := factorial(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
