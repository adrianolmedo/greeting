package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
	want := "cba"
	if got := ReverseRunes("abc"); got != want {
		t.Errorf("want: %v, got: %v", got, want)
	}
}
