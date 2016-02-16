package xstrings

import "testing"

func TestSubstrFromBeginning(t *testing.T) {
	commitID := "e3ff9321271b0a5cec45ca6e0cdc72b2f376afd2"
	expected := "e3ff9"
	if s := Substr(commitID, 0, 5); s != expected {
		t.Errorf("is %v, expect %v", s, expected)
	}
}

func TestSubstrFromCharFive(t *testing.T) {
	commitID := "e3ff9321271b0a5cec45ca6e0cdc72b2f376afd2"
	expected := "32127"
	if s := Substr(commitID, 5, 5); s != expected {
		t.Errorf("is %v, expect %v", s, expected)
	}
}
