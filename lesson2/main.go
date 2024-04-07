package main

import "fmt"

var (
	vrune rune
	vbyte byte
	var1  int = 16
	var2  int = 3
)

func main() {
	fmt.Println("Task2.1")
	fmt.Println("rune:", vrune, ", byte:", vbyte)
	fmt.Printf("rune: %c, byte: %q\n", vrune, vbyte)
	fmt.Println("потому что оба типа являются целочисленными и значением по умолчанию является 0")
	fmt.Println("Task2.2")
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %T\n", var1/var2, var1%var2, var1/var2)
}
