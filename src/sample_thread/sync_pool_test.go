package sample_thread

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

//对象缓存
//gc 会清空sync.pool 缓存对象
//总结
/*
适合于通过复用，降低复杂对象的创建和GC代价
协程安全，会有锁的开销
生命周期受GC影响，不适合于做连接池等，需自己管理生命周期的资源的池化
*/

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}
	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	//gc会清空 pool.Put(3)的值，
	runtime.GC()
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
	v2, _ := pool.Get().(int)
	fmt.Println(v2)
}

func TestSyncPoolMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
