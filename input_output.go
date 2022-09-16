package main

import (
	"fmt"
	"strconv"
)

func main() {
	var temp string

	fmt.Print("Enter First Number: ")
	fmt.Scanln(&temp)
	a, _ := strconv.Atoi(temp)

	fmt.Print("Enter Second Number: ")
	fmt.Scanln(&temp)
	b, _ := strconv.Atoi(temp)

	fmt.Println("The Sum Is:", a+b)
}
