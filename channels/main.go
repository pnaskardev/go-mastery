package main

import (
	"fmt"
	"runtime"
	"sync"
)

func ping(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range ch {
		fmt.Println("PING", n)
		ch <- n + 1
	}
}

func pong(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range ch {
		fmt.Println("PONG", n)
		ch <- n + 1
	}
}

func main() {

	// ============================= ASSIGNMENT 1 ====================================
	// channel := make(chan int)
	// var wg sync.WaitGroup

	// wg.Add(2)
	// go ping(channel, &wg)
	// go pong(channel, &wg)

	// // start
	// channel <- 1

	// go func() {
	// 	wg.Wait()
	// 	close(channel)
	// }()
	// wg.Wait()

	// ============================= ASSIGNMENT 2 ==========================

	// var wg sync.WaitGroup

	// workers := 5

	// jobs := make(chan int, workers)
	// results := make(chan int, workers)

	// // PRODUCERS FAN OUT
	// // THIS will be a go routine to load the jobs channel with integers
	// go func() {
	// 	// This is blocking
	// 	for i := range 100 {
	// 		jobs <- i
	// 	}

	// 	// once the for loop is done it means all the elements are loaded
	// 	close(jobs)
	// }()

	// // WORKERS
	// // Goroutine to get whatever is in the jobs channel
	// // sqaure and put it in the results channel
	// // In range of workers cause I need multiple workers
	// for range workers {
	// 	wg.Go(func() {
	// 		// THis is blocking
	// 		// Once all jobs are done go ahead
	// 		for i := range jobs {
	// 			results <- i * i
	// 		}
	// 	})
	// }

	// // Coordinator - Safely CLosing Go routines
	// go func() {
	// 	wg.Wait()
	// 	close(results)
	// }()

	// // FAN IN
	// final_results := []int{}
	// for i := range results {
	// 	final_results = append(final_results, i)
	// }

	// fmt.Println(final_results)
	// fmt.Println(len(final_results))

	// ============================= ASSIGNMENT 3 =========================

	// c1 := make(chan string)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c1 <- "C1"
	// }()

	// select {
	// case res := <-c1:
	// 	fmt.Println("RES :- ", res)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println("timeout 1")
	// }

	// c2 := make(chan string)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	c2 <- "C2"
	// }()
	// select {
	// case res := <-c2:
	// 	fmt.Println("RES :- ", res)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("timeout 1")
	// }

	// ===================== ASSIGNMENT 4 ================

	var wg sync.WaitGroup

	workers := runtime.NumCPU()
	jobs1 := make(chan int, workers)
	jobs2 := make(chan int, workers)
	results := make(chan int, workers)

	// Producers
	go func() {

		// Produce to 2 different channels
		for i := 0; i < 1000000; i++ {
			jobs1 <- i
		}

		for i := 1000000; i <= 5000000; i++ {
			jobs2 <- i
		}

		close(jobs1)
		close(jobs2)

	}()

	// Workers
	wg.Go(func() {
		for val := range jobs1 {
			results <- val
		}
	})

	wg.Go(func() {
		for val := range jobs2 {
			results <- val
		}
	})

	// Coordinator
	go func() {
		wg.Wait()
		close(results)
	}()

	// Consumer
	final_results := []int{}
	for i := range results {
		final_results = append(final_results, i)
	}

	fmt.Println(final_results)

}
