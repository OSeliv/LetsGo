package main

import "fmt"

var S string = "Hello, Go!"

func main() {
	fmt.Println("Task4.1")
	hello(S)

	fmt.Println("Task4.2")
	task2_3 := func() { // 4.2 & 4.3
		fmt.Println(S)
	}
	task2_3()

	fmt.Println("Task4.3")
	hello3(task2_3)

	fmt.Println("Task4.4")
	task4 := hello4() // 4.4
	task4()

	fmt.Println("Task4.5")
	defer hello("\nзавершение работы") //4.5
	hello(S)

	fmt.Println("Task4.6")
	x, y, n := 0, 1, 23 //4.6
	fmt.Print(x, " ", y, " ")
	seqFibonacci(x, y, n-2)

}

func hello(str string) { // 4.1 & 4.5
	fmt.Println(str)
}

func hello3(foo func()) { // 4.3
	foo()
}

func hello4() func() { // 4.4
	return func() {
		fmt.Println(S)
	}
}

func seqFibonacci(x, y, n int) { //4.6
	if n <= 0 {
		return
	}
	fmt.Print(x+y, " ")
	seqFibonacci(y, x+y, n-1)
}
