package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	stop := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i

		go func() {
			defer wg.Done()
			for {
				select {
				case <-stop:
					fmt.Println("stop горутина: ", i)
					return
				default:
					time.Sleep(time.Second)
					fmt.Println("сложные вычисления горутины: ", i)
				}
			}
		}()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")

		close(stop)
	}()
	wg.Wait()
}
