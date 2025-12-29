package main

import (
	"fmt"

	"github.com/pnaskardev/go-mastery/Pointers/models"
)

func makePointer() *int {
	value := 69
	return &value
}

func main() {

	p := *makePointer()

	fmt.Println(p)

	u := models.CreateNewUser()

	u.BirthdayByValue()

	fmt.Printf("BY VALUE - %+v\n", u)

	u.BirthdayByRef()

	fmt.Printf("BY REF - %+v\n", u)

	// user_nil := models.User{}

	error := models.UserNilCheck(nil)

	fmt.Println(error.Error())

}

// a, b := 5, 6

// 	fmt.Println("BEFORE SWAP", a, b)

// 	utils.Swap(&a, &b)

// 	fmt.Println("AFTER SWAP", a, b)
