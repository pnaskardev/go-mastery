package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	wg.Go(func() {
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	})

	wg.Go(func() {
		for range 1000 {
			mu.Lock()
			counter++
			mu.Unlock()
		}
	})

	wg.Wait()
	fmt.Println(counter)

}
