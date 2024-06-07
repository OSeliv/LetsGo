package main

import "fmt"

//-------- 12.1 -----------
func do(v any) {
	a := 10
	// begin code
	t, ok := v.(int)
	if ok {
		a += t
	} else {
		fmt.Println("нельзя привести к int")
		return
	}
	// end code
	fmt.Println(a)
}

// ======= 12.2 ========
type Bird interface {
	Fly()
}

type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}

func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}

func Do(b Bird) {
	b.Fly()
	//здесь нужно дописать код
	t, ok := b.(Duck)
	if ok {
		t.Swim()
	}
}

/// -------- 12.3 ---------
type Xml struct{}

func (x Xml) Format() {
	fmt.Println("Данные в формате xml")
}

type Csv struct{}

func (c Csv) Format() {
	fmt.Println("Данные в формате csv")
}

type Format interface { // add
	Format()
}

func Report(f Format) {
	f.Format()
}

// ==============

func main() {
	fmt.Println("Task12.1")
	a := 1 // a := "5"
	do(a)

	fmt.Println("Task12.2")
	var d, s Bird
	d = Duck{}
	Do(d)
	s = Sparrow{}
	Do(s)

	fmt.Println("Task12.3")
	x := Xml{}
	Report(x)
	c := Csv{}
	Report(c)
}
