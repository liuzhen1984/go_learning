package collection

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(s0, len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(s0, len(s0), cap(s0))

	s1 := []int{1, 2, 3, 4}
	t.Log(s1, len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(s2, len(s2), cap(s2))
	//t.Log(s2[0],s2[1],s2[2],s2[3],s2[4])
	t.Log(s2[0], s2[1], s2[2])

	s2 = append(s2, 1)
	t.Log(s2[0], s2[1], s2[2], s2[3])
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Now", "Dec"}
	Q1 := year[0:3]
	t.Log(Q1, len(Q1), cap(Q1))
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "Unknown"
	t.Log(Q2, len(Q2), cap(Q2))
	t.Log(summer, len(summer), cap(summer))
}

func TestSliceCompare(t *testing.T) {
	s1 := []int{1, 3, 4}
	s2 := []int{1, 3, 4}
	//if s1==s2 { // (slice can only be compared to nil)
	t.Log("s1==s2", s1, s2)
	//}
}
