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
