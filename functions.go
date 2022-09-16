package main

import "fmt"

func sum(a float32, b int)string {
	fmt.Println(a+float32(b))
	return "Hello"
}

func main() {
	fmt.Println(sum(15.5, 13))
}