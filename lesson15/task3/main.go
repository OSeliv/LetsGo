package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			n++
		}()
	}
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(n)
}
