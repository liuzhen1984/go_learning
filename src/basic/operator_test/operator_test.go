package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1,3,4,5,6}
	d := [...]int{1, 3, 4, 5}
	t.Log(a == b)
	//t.Log(a==c)
	t.Log(a == d)
	t.Log(d == b)

}

func TestBitClear(t *testing.T) {
	a := 7
	a = a &^ Readable
	a = a &^ Writable
	t.Log(a&Readable == Readable, Readable, a&Writable == Writable, Writable, a&Executable == Executable, Executable)

}
