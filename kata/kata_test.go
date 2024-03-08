package kata

import "testing"

func doTest(t *testing.T, got, expected int) {
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
