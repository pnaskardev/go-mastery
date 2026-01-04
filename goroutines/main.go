package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func heartBeat(ctx context.Context) {
	// TICKER IS A CHANNL
	ticker := time.NewTicker(500 * time.Millisecond)

	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("PULSE")
		case <-ctx.Done():
			fmt.Println("HEARTBEAT STOPPED")
			return
		}
	}
}

func main() {

	var global_counter int64 = 0

	start := time.Now()
	var wg1 sync.WaitGroup
	for range 10000 {
		wg1.Go(func() {
			fmt.Println("DONE")
			atomic.AddInt64(&global_counter, 1)
		})

	}

	wg1.Wait()
	elapsed := time.Since(start)
	fmt.Println("EXECUTION TIME:- ", elapsed)
	fmt.Println("GLOBAL COUNTER:-", global_counter)

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Go(func() {
			fmt.Println(i)
		})
	}

	wg.Wait()

	ctx, cancel := context.WithCancel(context.Background())

	go heartBeat(ctx)

	time.Sleep(5 * time.Second) // simulate work

	cancel()

	time.Sleep(200 * time.Millisecond)

}
