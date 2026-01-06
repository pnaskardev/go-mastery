package main

import (
	"context"
	"fmt"
	"time"
)

func slowAPICall(ctx context.Context) error {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("WORK DONE")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func worker(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		fmt.Println(name, "cancelled:", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println(name, "finished work")
	}
}

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	start := time.Now()
	// call the slowApiCall
	// Context with cancel
	err := slowAPICall(ctx)
	end := time.Since(start)

	fmt.Println("ERROR:", err)
	fmt.Println("TIME ELAPSED", end)

	parentContext, parentCancel := context.WithCancel(context.Background())
	childContext1, _ := context.WithCancel(parentContext)
	childContext2, _ := context.WithTimeout(parentContext, 10*time.Second)

	go worker(childContext1, "child-1")
	go worker(childContext2, "child-2")

	time.Sleep(1 * time.Second)

	fmt.Println("Cancelling parent context...")
	parentCancel()
	time.Sleep(500 * time.Millisecond)

}
