package sample_thread

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//注意goroutine 泄漏问题
func runTask(i int) string {
	time.Sleep(time.Microsecond * time.Duration(i))
	return fmt.Sprintf("runTask %d", i)
}

func FirstResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
