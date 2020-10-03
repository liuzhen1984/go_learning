package exception

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("Something wrong!"))
}
