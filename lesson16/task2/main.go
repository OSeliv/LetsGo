package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		i := i
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("отмена горутины: ", i, ctx.Err())
					return
				default:
					fmt.Println("работа горутины: ", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("стоп!")
		cancel()
	}()
	wg.Wait()
}
