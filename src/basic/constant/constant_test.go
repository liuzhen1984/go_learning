package constant

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}
func TestConstantTry1(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, Readable, a&Writable == Writable, Writable, a&Executable == Executable, Executable)
}
