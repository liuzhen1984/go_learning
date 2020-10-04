package threads

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("start service")
	time.Sleep(time.Second * 1)
	return "Done"
}
func otherService() {
	fmt.Println("start otherService")
	time.Sleep(time.Second * 1)
	fmt.Println("end otherService")
}

func TestSyncService(t *testing.T) {
	t.Log(service())
	otherService()
	t.Log("end")
}

func AsyncService() chan string {
	//buffer channel 就需要等到 fmt.println执行完才执行service.exited
	//retCh := make(chan string)
	//buffer channel 就不需要等到 fmt.println执行完才执行service.exited
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service.exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherService()
	fmt.Println(<-retCh)
}
