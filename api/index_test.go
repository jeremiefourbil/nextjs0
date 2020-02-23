package api

import "testing"

func TestSum(t *testing.T) {
	total := Sum(4, 4)
	if total != 8 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}