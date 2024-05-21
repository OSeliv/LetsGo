package main

import (
	"fmt"
	"strings"
)

var (
	v1, v2, v3, v4, v5, v6  int8 = 1, 2, 3, 5, 10, 20
	v11, v12, v13, v14, v15 int  = 1, 2, 3, 5, 10
)

type square int // 5.6

func main() {
	var p *string // 5.1 & 5.3

	var s string = "Hello, Kitty!" // 5.2  & 5.3 & 5.5

	var s6 square = 25 //5.6 I'm sorry, but "s" has used

	// --- this block was checked---
	fmt.Println("Task 5.2")
	fmt.Printf("s значение: %s, адрес: %p\n", s, &s) //5.2

	fmt.Println("Task 5.3")
	p = &s                                         // 5.3
	*p = strings.Replace(*p, "Hello", "Goodby", 1) //changing var s
	fmt.Printf("s значение: %s\n", s)              // this is a magic

	fmt.Println("Task 5.4")
	fmt.Printf("v1: %d, %p\n", v1, &v1) // 5.4 1st group
	fmt.Printf("v2: %d, %p\n", v2, &v2) // &v1 + 1b(size of int8)
	fmt.Printf("v3: %d, %p\n", v3, &v3) // &v2 + 1b
	fmt.Printf("v4: %d, %p\n", v4, &v4) // &v3 + 1b
	fmt.Printf("v5: %d, %p\n", v5, &v5) // &v4 + 1b
	fmt.Printf("v6: %d, %p\n", v6, &v6) // &v5 + 1b

	fmt.Printf("v11: %d, %p\n", v11, &v11) // 5.4 2nd group
	fmt.Printf("v12: %d, %p\n", v12, &v12) // &v11 + 8b(size of int64)
	fmt.Printf("v13: %d, %p\n", v13, &v13) // &v12 + 8b
	fmt.Printf("v14: %d, %p\n", v14, &v14) // &v13 + 8b
	fmt.Printf("v15: %d, %p\n", v15, &v15) // &v14 + 8b
	// ---- end of the block  ----

	fmt.Println("Task 5.5")
	change(&s)                        //to change var s
	fmt.Printf("s значение: %s\n", s) // it's new

	fmt.Println("Task 5.6")
	fmt.Println(s6) //square type

	fmt.Println("Task 5.7")
	var s7 square = 30 //5.7 I'm sorry, but "s" has used for string
	s7 += 15           // changing value
	fmt.Println(s7)

	fmt.Println("Task 5.8")
	var s8 square = 34         //5.8
	s8 += 10                   // changing value
	fmt.Println(squareOut(s8)) // I'm not sure this has been the right solution

}

func change(p *string) { // 5.5
	*p = "GG WP" //changing param
}

func squareOut(sq square) string { //5.8
	return fmt.Sprintf("%d м²", sq)
}
