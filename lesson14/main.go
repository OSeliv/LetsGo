package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Task14.1")
	go func() {
		fmt.Println("Привет из дочерней горутины!")
	}()
	time.Sleep(2 * time.Second)

	fmt.Println("Task14.2")
	ch2 := make(chan string)
	go func() {
		ch2 <- "Привет, строковый канал!"
	}()

	fmt.Println(<-ch2)

	fmt.Println("Task14.3")
	ch3 := make(chan string, 4)
	go func() {
		ch3 <- "Привет"
		ch3 <- "буферизованный канал!"
	}()
	fmt.Println(<-ch3)
	fmt.Println(<-ch3)

	fmt.Println("Task14.4")
	ch4 := make(chan int)
	stop := make(chan struct{})
	close(ch4) // was added
	go func() {
		<-ch4
		stop <- struct{}{}
	}()
	<-stop
	fmt.Println("happy end")
}
