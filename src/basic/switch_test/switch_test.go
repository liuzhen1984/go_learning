package switch_test

import "testing"

func TestSwitchMultCase(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3, 5:
			t.Log("Odd")
		default:
			t.Log("Other value")

		}
	}
}

func TestSwitchCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")

		}
	}
}
