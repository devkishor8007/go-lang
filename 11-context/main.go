package main

import (
	"fmt"
	"context"
)

func tryGo(ctx context.Context) {
	fmt.Println("learning GO")
}

func newOrder(ctx context.Context) {
	fmt.Printf("do something is %s\n", ctx.Value("meroKey"))
}

func main() {
  ctx := context.Background()

  ctx = context.WithValue(ctx, "meroKey", "hello the value of key is thisValue")
  tryGo(ctx)
  newOrder(ctx)
}

// we can use - context.Background() or context.TODO()
// both function creates an empty context
// tips: context.Background is a good default option.