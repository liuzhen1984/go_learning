package threads

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}
func cancel2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestRunMultipleGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	cancelChan := make(chan struct{})
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}, wg *sync.WaitGroup) {
			wg.Add(1)
			for {
				if isCancelled(cancelChan) {
					wg.Done()
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, cancelChan, &wg)
	}
	//cancel1 only close one routine
	//cancel1(cancelChan)
	time.Sleep(time.Second * 1)
	cancel2(cancelChan)
	wg.Wait()
}
