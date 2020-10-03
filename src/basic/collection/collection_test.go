package collection

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 56}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idx, e := range arr3 {
		t.Log(idx, e)
	}

	for _, e := range arr3 {
		t.Log(e)
	}
	t.Log("=========split array=========")
	t.Log(arr3[1:2])
	t.Log(arr3[1:3])
	t.Log(arr3[1:len(arr3)])
	t.Log(arr3[1:])
	t.Log(arr3[:3])
	x := arr3[1:2]
	y := arr3[1:3]
	t.Log(x, y)
	x[0] = 10
	t.Log(x, y, arr3)
}
