package kata

import "testing"

func doTest(t *testing.T, got, expected int) {
	if expected != got {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestConvertIpToInt(t *testing.T) {
	doTest(t, ConvertIpToInt("0.0.0.0"), 0)
	doTest(t, ConvertIpToInt("0.0.0.1"), 1)
	doTest(t, ConvertIpToInt("0.0.1.1"), 257)
	doTest(t, ConvertIpToInt("0.1.2.1"), 256 * 256 + 256 * 2 + 1)
}
