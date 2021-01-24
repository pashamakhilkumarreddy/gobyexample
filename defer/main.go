package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defer example")
	defer printMessage()
	fmt.Println("Uno")
	fmt.Println("Dos")
	fmt.Println("Tres")
}

func printMessage() {
	fmt.Println("Defer function execution")
}
