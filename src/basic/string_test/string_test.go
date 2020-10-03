package string_test

import (
	"strconv"
	"strings"
	"testing"
)

//1.String is a byte slice of only reading
//2.the byte slice can save any type data

func TestString(t *testing.T) {
	var s string
	t.Logf("s='%s'", s)
	s = "hello"
	t.Logf("s='%s',len=%d", s, len(s))
	s = "\xE4\xB8\xA5"
	t.Logf("s='%s',len=%d", s, len(s))
	s = "中"
	//s[1]='3' //cannot assign to s[1]
	t.Logf("s='%s',len=%d", s, len(s))
	// rune is an alias for int32 and is equivalent to int32 in all ways. It is
	// used, by convention, to distinguish character values from integer values.
	c := []rune(s)                          //rune can get
	t.Logf("中 unicode %x", c[0])            //unicode 是一个字符集(code point)
	t.Logf("中 UTF8 %x,unicode %x", s, c[0]) //UTF8 是unicode的存储实现（转换为字节序列的规则）
	//strings (https://golang.org/pkg/strings/)
	//strconv (https://golang.org/pkg/strconv/)
}

func TestTravelString(t *testing.T) {
	s := "中华人民共共和国"
	for idx, v := range s {
		t.Logf("%[1]c %[1]x index=%d, ", v, idx)
		t.Logf("index=%d %[1]c %[1]x , ", idx, v) //format is wrong
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(strings.Join(parts, "-"))
}

func TestStringCovn(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str=" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
