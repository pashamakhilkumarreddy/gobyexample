package main

import (
	"fmt"
)

func main() {
	fmt.Println("Panic Example")
	defer fmt.Println("This is executed even when panic is called (before panic executes)")
	executePanic()
	fmt.Println("Main block executes completely")
}

func executePanic() {
	panic("Panic starts")
	fmt.Println("Function executes completely")
}
