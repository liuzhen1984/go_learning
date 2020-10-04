package sample_thread

//可使用sync.WaitGroup完成
//使用另外一种方式

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

//注意goroutine 泄漏问题
func runTask1(i int) string {
	time.Sleep(time.Microsecond * time.Duration(i))
	return fmt.Sprintf("runTask %d", i)
}

func AllResponse() string {
	numOfRunner := 5
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask1(i)
			ch <- ret
		}(i)
	}
	finalRet := "\n"
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestFirstResponse1(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())

}
