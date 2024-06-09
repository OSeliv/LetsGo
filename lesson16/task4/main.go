package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {
	ctx := context.Background()
	var key1, key2 ctxKey = "some key1", "some key2"
	ctx = context.WithValue(ctx, key1, "some value1")
	ctx = context.WithValue(ctx, key2, "some value2")

	do(ctx)
}

func do(ctx context.Context) {
	var key1, key2 ctxKey = "some key1", "some key2"
	fmt.Println(ctx.Value(key1))
	fmt.Println(ctx.Value(key2))
}
