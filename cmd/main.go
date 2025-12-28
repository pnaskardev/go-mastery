package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pnaskardev/calc/basic"
	"github.com/pnaskardev/calc/utils"
	mylogger "github.com/pnaskardev/my-logger"
)

func main() {

	// Topic A: Go Mod, Tooling & Workspace
	fmt.Println(basic.Add(2, 3))
	fmt.Println(basic.Subtract(2, 3))

	fmt.Println(uuid.NewString())

	// logrus.Info("HELLO")

	mylogger.Info("TEST")

	// Topic B: Pointers & Memory Mechanics

	a, b := 5, 6

	fmt.Println("BEFORE SWAP", a, b)

	utils.Swap(&a, &b)

	fmt.Println("AFTER SWAP", a, b)
}
