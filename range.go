// Range is like Python Enumerate
package main

import (
	"fmt"
)

func main() {
	a := "Hello"
	for i, num := range a {
		fmt.Println(i, string(num))
	}
}