package extension

import (
	"fmt"
	"testing"
)

type Pet2 struct {
}

func (p *Pet2) Speak() {
	fmt.Print("....")
}

func (p *Pet2) SpeekTo(host string) {
	p.Speak()
	fmt.Print(" ", host)
}

type Dog2 struct {
	Pet2
}

func (d *Dog2) Speak() {
	fmt.Print("Wang!")
}

func TestDog2(t *testing.T) {
	dog := new(Dog2)
	//p:=Pet2(dog)//cannot convert dog (type *Dog2) to type Pet2
	dog.SpeekTo("dog")
}
