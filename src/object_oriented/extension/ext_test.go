package extension

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("....")
}

func (p *Pet) SpeekTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

//func (d *Dog) Speak(){
//	d.p.Speak()
//}
//func (d *Dog) SpeekTo(host string) {
//	d.p.SpeekTo(host)
//}
func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func (d *Dog) SpeekTo(host string) {
	d.Speak()
	fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeekTo("dog")
}
