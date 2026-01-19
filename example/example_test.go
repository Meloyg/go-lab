package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 6 {
		t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test add succ")
}

func TestSub(t *testing.T) {
	r := sub(10, 2)

	if r != 8 {
		t.Fatalf("Wrong")
	}

	t.Logf("test sub succ")
}
