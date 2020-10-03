package loop

import "testing"

func TestWhileLoop(t *testing.T) {
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

func TestUnlimitLoop(t *testing.T) {
	n := 0
	for {
		t.Log(n)
		n++
		if n > 10 {
			break
		}
	}
}
