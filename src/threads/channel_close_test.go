package threads

import (
	"fmt"
	"sync"
	"testing"
)

func doProducer(ch chan int, wd *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wd.Done()
	}()

}

func doReceiver(ch chan int, wd *sync.WaitGroup) {
	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wd.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wd sync.WaitGroup
	ch := make(chan int, 10)
	wd.Add(1)
	doProducer(ch, &wd)
	wd.Add(1)
	doReceiver(ch, &wd)
	wd.Add(1)
	doReceiver(ch, &wd)
	wd.Add(1)
	doReceiver(ch, &wd)
	wd.Add(1)
	doReceiver(ch, &wd)
	wd.Wait()

}
