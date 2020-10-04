package threads

import (
	"fmt"
	"testing"
	"time"
)

func service1() string {
	fmt.Println("start service")
	time.Sleep(time.Second * 2)
	return "Done"
}
func otherService1() string {
	fmt.Println("start otherService")
	time.Sleep(time.Second * 2)
	return "otherDown"
}

func AsyncService1() chan string {
	//buffer channel 就需要等到 fmt.println执行完才执行service.exited
	//retCh := make(chan string)
	//buffer channel 就不需要等到 fmt.println执行完才执行service.exited
	retCh := make(chan string)
	go func() {
		ret := service1()
		retCh <- ret
	}()
	return retCh
}
func AsyncOtherService1() chan string {
	//buffer channel 就需要等到 fmt.println执行完才执行service.exited
	//retCh := make(chan string)
	//buffer channel 就不需要等到 fmt.println执行完才执行service.exited
	retCh := make(chan string, 2)
	go func() {
		ret := otherService1()
		retCh <- ret
	}()
	return retCh
}

func TestAsyncService1(t *testing.T) {

	select {
	case ret := <-AsyncService1():
		fmt.Printf("service %s", ret)
	case ret2 := <-AsyncOtherService1():
		fmt.Println("other service", ret2)
	case <-time.After(time.Second * 2):
		t.Error("timeout")

		/*

		  如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
		 否则：
		 如果有 default 子句，则执行该语句。如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。
		*/
		///*default:
		//t.Error("No one returned")
	}
}
