package calculate

import "testing"

func TestCal(t *testing.T) {
	result := Calculate(2)
	if result != 2*2 {
		t.Fatal("calculateの答えが違うよ。")
	}
}
