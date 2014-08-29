package tests

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(10, 20) != 30 {
		t.Error("expected value: 30")
	}
}
