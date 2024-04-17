package main

import "fmt"

const Pi = 3.14159 // 3.1 + 3.3

const ( // 3.4 + 3.7
	C1 = iota
	C2
	C3
	C4
	C5
)

func main() {
	const (
		pi float32 = 3.14 // 3.2
		n  int     = 5    // 3.6
	)

	fmt.Println("Task3.1")
	fmt.Println("Global Pi = ", Pi)

	fmt.Println("Task3.2")
	fmt.Println("local pi = ", pi, "Pi = ", Pi)

	fmt.Println("Task3.3")
	fmt.Printf("Global Pi was %f, local pi = %f\n", Pi, pi)
	const Pi = 3.1 // 3.3
	fmt.Printf("Global Pi became %f, local pi = %f\n", Pi, pi)

	fmt.Println("Task3.4")
	fmt.Printf("Global consts: %d, %d, %d, %d, %d\n", C1, C2, C3, C4, C5)

	fmt.Println("Task3.5 & Task3.7")
	const ( // 3.5 + 3.7
		c1 = 1 << iota
		c2
		c3
		c4
		c5
	)
	fmt.Printf("Local consts 2^n: %d, %d, %d, %d, %d\n", c1, c2, c3, c4, c5)

	fmt.Println("Task3.6")
	var m float32 = float32(n) + 3.4 // 3.6
	fmt.Println("var m = ", m)
}
