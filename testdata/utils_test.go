package testdata

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	if result != 5 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 5)
	}
}

func TestReverseWord(t *testing.T) {
	result := ReverseWord("planet")
	if result != "tenalp" {
		t.Errorf("Reverse was incorrect, got: %s, want: %s.", result, "tenalp")
	}
}