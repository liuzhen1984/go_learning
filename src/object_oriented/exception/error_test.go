package exception

import (
	"github.com/pkg/errors"
	"testing"
)

var LessThanTwoError = errors.New("n less 2]")
var MoreThanTwoError = errors.New("n larger 100]")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, MoreThanTwoError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}
func TestGetFibonacci(t *testing.T) {
	//the var life cycle untill  the else end
	if v, err := GetFibonacci(5); err != nil {
		if err == LessThanTwoError {
			t.Log("less")
		}
		if err == MoreThanTwoError {
			t.Log("more")
		}
	} else {
		t.Log(v)
	}
	if v, err := GetFibonacci(1); err == nil {
		t.Log(v)

	} else {
		if err == LessThanTwoError {
			t.Log("less")
		}
		if err == MoreThanTwoError {
			t.Log("more")
		}
	}
}
