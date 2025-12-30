package models

import (
	"fmt"
)

type Base struct {
	Age int
}

type Child struct {
	Base
	Age int
}

// func (c *Child) Describe() {
// 	fmt.Println("Child Describe", c.Age)
// }

func (b *Base) Describe() {
	fmt.Println("Base Describe", b.Age)
}

func NewChildInstance(childAge int, baseAge int) *Child {

	return &Child{
		Base: Base{
			Age: baseAge,
		},
		Age: childAge,
	}

}
