package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 9
	fmt.Println("Total is", sum(a, b))
	fmt.Println("Recursive is", recursive(c))
}

func sum(a int, b int) int {
	return a + b
}

func recursive(value int) int {
	if value > 1 {
		return value + recursive(value-1)
	} else {
		return value
	}
}
