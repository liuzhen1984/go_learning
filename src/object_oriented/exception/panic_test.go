package exception

import (
	"github.com/pkg/errors"
	"os"
	"testing"
)

//panic and os.Exit

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("defer ....")
	}()
	t.Log("processing...")
	panic(errors.New("test panic"))
}

//os.Exit will do not execute defer, does not println error stack data

func TestExit(t *testing.T) {
	defer func() {
		t.Log("defer ....")
	}()
	t.Log("processing...")
	os.Exit(-1)
}
