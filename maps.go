package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["Hell"] = -1
	m["Heaven"] = 1
	m["Purgatory"] = 0

	fmt.Println("Afterlife:", m, len(m))
	if val, ok := m["Purgator"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("Not Exists!")
	}
}
