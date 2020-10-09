package main

import (
	"fmt"
)

func main() {
	var name = "Stacy"
	var age uint32 = 25
	area := "Greenwich"
	var city string
	city = "London"
	var country string
	country = "England"
	fmt.Printf("I am %v of age %d from %s, %s, %s. \n", name, age, area, city, country)
}
