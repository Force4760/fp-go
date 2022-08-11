package fp

import (
	"testing"
)

func TestId(t *testing.T) {
	res := Id(42)
	if res != 42 {
		t.Error("Id should have returned 42. Received:", res)
	}
}

func TestIf_True(t *testing.T) {
	res := If(true, 42, 21)
	if res != 42 {
		t.Error("If should have returned 42. Received:", res)
	}
}
func TestIf_False(t *testing.T) {
	res := If(false, 21, 42)
	if res != 42 {
		t.Error("If should have returned 42. Received:", res)
	}
}
