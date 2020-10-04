package threads

import (
	"sync"
	"testing"
	"time"
)

//有问题的
func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}
func TestCounterSafe(t *testing.T) {
	/**
	// A Mutex is a mutual exclusion lock.
	// The zero value for a Mutex is an unlocked mutex.
	//
	// A Mutex must not be copied after first use.
	*/
	var mut sync.Mutex //互斥锁
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	//这里可以使用sync.waitGroup实现
	time.Sleep(time.Second * 1)
	t.Logf("counter = %d", counter)
}

func TestCounteSafe(t *testing.T) {
	/**
	// A Mutex is a mutual exclusion lock.
	// The zero value for a Mutex is an unlocked mutex.
	//
	// A Mutex must not be copied after first use.
	*/
	var mut sync.Mutex //互斥锁
	//A WaitGroup waits for a collection of goroutines to finish.
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			defer func() {
				wg.Done()
			}()
			mut.Lock()
			counter++
		}()
	}
	wg.Wait()
	t.Logf("counter = %d", counter)
}
