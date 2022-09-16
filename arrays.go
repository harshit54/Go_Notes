package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a [2][2]int
	fmt.Println("Original:", a)
	var temp string
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scanln(&temp)
			a[i][j], _ = strconv.Atoi(temp)
		}
	}
	fmt.Println("Final:", a)
}
