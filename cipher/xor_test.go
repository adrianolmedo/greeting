package cipher

import "testing"

func TestXor(t *testing.T) {
	target := "Hello world"
	if got := Xor(target); got == target {
		t.Errorf("Xor(\"Hello world\") ==> got: %q", got)
	}
}
