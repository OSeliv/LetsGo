package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Task14.6")
	explanation := `
	It's very close to Task5: 
	2 child goroutines - the 1st reads values from the channel _ch_, the 2nd sends values to _ch_ once per second. 
	And the main goroutine, that waits for 5 sec and then send 2 stop-signal to terminate child goroutines. 
	Therefore, child goroutines work for a certain duration. 
	Main goroutine waits for another sec then finishes with printing message.`
	fmt.Println(explanation)
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	// 1st child goroutine
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	// 2nd child gorotine
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
