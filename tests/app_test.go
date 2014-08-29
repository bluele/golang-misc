package tests

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(10, 20) != 30 {
		t.Error("expected value: 30")
	}
}

func TestMul(t *testing.T) {
	if Mul(10, 20) != 200 {
		t.Error("expected value: 200")
	}
}
