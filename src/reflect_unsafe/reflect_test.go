package reflect_unsafe

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"testing"
)

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestCheckType(t *testing.T) {
	var f int64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
	var f1 float32 = 1
	CheckType(&f1)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"` //struct tag
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

//两个不同struct之间，相同field 值的传递和赋值
func fillBySettings(s interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(s).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer")
	}
	if reflect.TypeOf(s).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}
	if settings == nil {
		return errors.New("settings is nil.")
	}
	var fields reflect.StructField
	var ok bool

	for k, v := range settings {
		if fields, ok = (reflect.ValueOf(s)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if fields.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(s)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFill(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 10}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)

	v := 10
	if err := fillBySettings(&v, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Logf("Name: value(%[1]v), Type(%[1]T) ", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get 'Name' field.")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(31)})
	t.Log("Updated Age:", e)
}

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 3: "three", 2: "two"}
	//t.Log(a == b) //invalid operation: a == b (map can only be compared to nil)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1, 3}
	t.Log("s1==s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1==s3?", reflect.DeepEqual(s1, s3))
}
