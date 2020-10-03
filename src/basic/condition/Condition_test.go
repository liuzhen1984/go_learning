package condition

import "testing"

func TestIfMultiSec(t *testing.T) {
	if v, err := someFun(); err == "" {
		t.Log("1==1", v)
	} else {
		t.Log("1!=1")
	}
}
func someFun() (int8, string) {
	return 2, ""
}
