package threads

//取消任务和子任务

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func isCancelled1(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancelContextGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context, wg *sync.WaitGroup) {
			wg.Add(1)
			for {
				if isCancelled1(ctx) {
					wg.Done()
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, ctx, &wg)
	}

	time.Sleep(time.Second * 1)
	cancel()
	wg.Wait()
}
