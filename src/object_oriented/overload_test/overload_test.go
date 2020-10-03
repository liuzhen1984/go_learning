package overload_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(hello world)"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(Hello World)"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld()) //%T print object type
}

func TestPolymorphism(t *testing.T) {
	goProg := &GoProgrammer{}
	//goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}

///empty interface

func CheckType(p interface{}) {
	//if i,ok := p.(int); ok{
	//	fmt.Println("integer ",i)
	//	return
	//}
	//if i,ok :=p.(string);ok{
	//	fmt.Println("string ",i)
	//	return
	//}
	switch v := p.(type) {
	case int:
		fmt.Println("integer ", v)
	case string:
		fmt.Println("string ", v)
	default:
		fmt.Println("unknown ", v)

	}
}

func TestCheckType(t *testing.T) {
	CheckType(10)
	CheckType("String")
	CheckType(new(GoProgrammer))

}
