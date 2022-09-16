package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("Init:", s, len(s))
	s[0] = "Just"
	s[1] = "Do"
	s[2] = "It!"
	fmt.Println("Final:", s, len(s))
	s = append(s, "By", "Nike", "✔️  ")
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("Reversed:", s)
}