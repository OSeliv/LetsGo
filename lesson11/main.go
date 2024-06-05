package main

import (
	"errors"
	"fmt"
)

// --- 11.4 ---
type myFirstError struct {
	message string
}

func (e *myFirstError) Error() string {
	return e.message
}

// ------------

func main() {
	fmt.Println("Task11.1")
	//11.1 + 11.2 + 11.3 + 11.4
	err := errors.New("ошибка1")
	err1 := fmt.Errorf("ошибка3:%w", fmt.Errorf("ошибка2:%w", err))
	fmt.Println(err1)

	fmt.Println("Task11.2")
	//errchain from 11.1
	err2 := errors.Unwrap(err1)
	fmt.Println(err2)

	fmt.Println("Task11.3")
	//errchain from 11.1 + err = ошибка1
	fmt.Println(err1)
	fmt.Printf("\"%v\" была: %v\n", err, errors.Is(err1, err))

	fmt.Println("Task11.4")
	me := &myFirstError{message: "ошибка1"}
	//errchain from 11.1
	fmt.Println(err1)
	fmt.Println("ошибки типа myFirstError не было:", !errors.As(err1, &me)) //сhecking NOT As
}
