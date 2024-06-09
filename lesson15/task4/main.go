package main

import (
	"fmt"
	"sync"
)

var loadOnce sync.Once

func start() {
	fmt.Println("starting...")
}
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			loadOnce.Do(start)
			fmt.Println("goroutine", j)
		}()
	}
	wg.Wait()
}
