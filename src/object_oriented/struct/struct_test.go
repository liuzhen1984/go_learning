package _struct_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	te := e
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	te2 := e2
	t.Log(e, te)
	te.Id = "10"
	t.Log(e, te)

	t.Log(e2, te2)
	te2.Id = "100"
	t.Log(e2, te2)

	t.Log(e1)
	t.Log(e1.Id)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

//same address
//func (e *Employee) String() string{
//	fmt.Printf("address is %x\n",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("ID:%s/Name:%s/Age:%d",e.Id,e.Name,e.Age)
//}
//different address
func (e Employee) String() string {
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}
func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
func TestStructOperationsPoint(t *testing.T) {
	e := &Employee{"0", "Bob", 20}
	fmt.Printf("address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}
