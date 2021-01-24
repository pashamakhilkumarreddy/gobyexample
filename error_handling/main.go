package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Error Handling in Golang")
	var (
		a int = 16
		b int = 0
	)
	result, err := divide(a, b)
	if err != nil {
		err := fmt.Errorf("Error is: %v", err)
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Result of %v / %v is %v", a, b, result)
}

func divide(a, b int) (float64, error) {
	if b == 0 {
		return 0, errors.New("Denominator must be greater than 0")
	}
	return float64(a / b), nil
}
