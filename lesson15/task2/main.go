package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var j int32
	for i := 0; i < 1000; i++ {
		go func() {
			atomic.AddInt32(&j, 1)
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Println(j)
}
