package main

import (
	"fmt"
)

func makePointer() *int {
	value := 69
	return &value
}

func main() {

	p := *makePointer()

	fmt.Println(p)

}

// a, b := 5, 6

// 	fmt.Println("BEFORE SWAP", a, b)

// 	utils.Swap(&a, &b)

// 	fmt.Println("AFTER SWAP", a, b)
