package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("отмена горутина: ", i)
					return
				default:
					fmt.Println("работа горутины: ", i)
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
