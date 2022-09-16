package main

import "fmt"

func arraySum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	arraySum(1, 2, 3)
	arraySum(1, 2, 3, 4, 5)
	arraySum(arr...)
}
