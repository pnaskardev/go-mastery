package models

import "fmt"

type Drivable struct {
}

func (d *Drivable) Drive() {
	fmt.Println("DRIVING")
}

type Flyable struct {
}

func (d *Flyable) Fly() {
	fmt.Println("FLYING")
}

type FlyingCar struct {
	Drivable
	Flyable
}
