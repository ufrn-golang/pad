package pad

import (
	"testing"
	"testing/quick"
)

func TestPad(t *testing.T) {
	length := 8
	if result := Pad("test", uint(length)); len(result) != length {
		t.Errorf("Expected %d, got %d", length, len(result))
	}
}

func TestPadGenerative(t *testing.T) {
	isEqual := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}
	if err := quick.Check(isEqual, nil); err != nil {
		t.Error(err)
	}
}