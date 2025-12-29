package models

import "fmt"

type User struct {
	Name     string
	Age      int
	Birthday string
}

var UserInstance User

func CreateNewUser() User {

	UserInstance = User{
		Name:     "Priyanshu Naskar",
		Age:      23,
		Birthday: "28 march",
	}

	return UserInstance
}

func (u User) BirthdayByValue() {
	u.Age++
}

func (u *User) BirthdayByRef() {
	u.Age++
}

func UserNilCheck(u *User) error {
	if u == nil {
		return fmt.Errorf("u_is_nil")
	}

	return nil
}
