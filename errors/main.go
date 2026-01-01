package main

import (
	"errors"
	"fmt"
)

func Divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("DENOMINATOR_IS_0")
	}

	val := a / b
	return val, nil
}

var ErrorNotFound = errors.New("ERROR_USER_NOT_FOUND")

func FindUserByID(id int) (string, error) {

	if id == 1 {
		return "user_found", nil
	}

	return "user_not_found", ErrorNotFound
}

type AppError struct {
	Code    int
	Message string
}

func (err *AppError) Error() string {
	return fmt.Sprintf("Code %d: %s", err.Code, err.Message)
}

func someFunction() error {
	// Return a wrapped custom error
	// return fmt.Errorf("operation failed: %w", &AppError{Code: 101, Message: "not found"})

	return &AppError{Code: 101, Message: "not found"}
}

func main() {

	// result, err := Divide(2, 0)
	// if err != nil {
	// 	fmt.Println("ERROR -", err.Error())
	// 	return
	// }

	// fmt.Printf("RESULT - %+v\n", result)
	// ============================================================
	// user_find_result, find_user_error := FindUserByID(69)

	// if find_user_error != nil {
	// 	if errors.Is(find_user_error, ErrorNotFound) {
	// 		fmt.Println("USER_NOT_FOUND", find_user_error)
	// 		return
	// 	}
	// 	fmt.Println("SOMETHING WENT WRONG")
	// } else {
	// 	fmt.Printf("%+v", user_find_result)
	// }

	// ==============================================================

	err := someFunction()

	var customErr *AppError

	if errors.As(err, &customErr) {
		fmt.Printf("Caught custom error AppError. Code: %d, Message: %s\n", customErr.Code, customErr.Message)
		fmt.Println(customErr.Error())
	} else {
		fmt.Println("Error is not of type *AppError")
	}
}
