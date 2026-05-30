package main

import (
	"testing"
)

func TestMaxInteger(t *testing.T) {
	a, b := 2, 7

	res := MaxInteger(a, b)

	if res != b {
		t.Errorf("expected %d, got %d", b, res)
	}
}

func TestMain(m *testing.M) {
	main()
}
