package main

import (
	"fmt"
)

type Contract struct { //6.1, 6.3
	ID           int
	Number, Date string
}

func (c Contract) String() string { // 6.3
	return fmt.Sprintf("Договор № %s от %s", c.Number, c.Date)
}

// --- 6.4 ---
type contact struct {
	Addss string
	Phone string
}

type user struct {
	ID   int
	Name string
	contact
}

type employee struct {
	ID   int
	Name string
	contact
}

// ----

func main() {

	fmt.Println("Task 6.1")                        //6.1
	c1 := Contract{1, "#000A\\n101", "2024-01-31"} //&6.3
	fmt.Printf("%+v\n", c1)

	fmt.Println("Task 6.2") //6.2
	type contract struct {
		ID           int
		Number, Date string
	}
	c2 := contract{1, "#000A101\t01", "2024-01-31"}
	fmt.Printf("%+v\n", c2)

	fmt.Println("Task 6.3") //6.3
	fmt.Println(c1)

	fmt.Println("Task 6.4") //6.4
	u := user{
		ID:   2,
		Name: "Queen Elizabeth",
		contact: contact{
			"Buckingham Palace",
			"Buckingham-2",
		},
	}
	e := employee{
		ID:   7,
		Name: "James Bond",
		contact: contact{
			"Casino Royale",
			"agent-007",
		},
	}
	fmt.Println("user:", u.Name, "\t\temployee:", e.Name)
	fmt.Println(u.Addss, u.Phone, "\t", e.Addss, e.Phone)
}
