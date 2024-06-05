package main

import (
	"first"  //10.1, 10.3
	"second" //10.3

	"fmt"
)

func main() {
	fmt.Println("Task 10.1")
	f := first.Hello()
	fmt.Println(f)

	fmt.Println("Task 10.3")
	fmt.Println(first.Hello(), second.Hello())
}
