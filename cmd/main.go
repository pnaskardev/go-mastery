package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pnaskardev/calc/basic"

	mylogger "github.com/pnaskardev/my-logger"
)

func main() {

	// Topic A: Go Mod, Tooling & Workspace
	fmt.Println(basic.Add(2, 3))
	fmt.Println(basic.Subtract(2, 3))

	fmt.Println(uuid.NewString())

	// logrus.Info("HELLO")

	mylogger.Info("TEST")

}
