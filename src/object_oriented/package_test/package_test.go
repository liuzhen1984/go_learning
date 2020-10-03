package package_test

import (
	"fmt"
	"testing"
)

// 小写开头不能被其他package 调用
func square(i int) int {
	return 0
}

//相同的方法在同一个package中不能重名

//init 方法 同一个包，同一个文件中可以定义多个，并且都会被执行

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func TestPackage(t *testing.T) {
	t.Log("start")
}

//go get -u github.com/golang....
