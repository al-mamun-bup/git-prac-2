package main

import "fmt"

func main() {
	fmt.Println("hello")

	fmt.Println(add(3, 4))
}

func add(x, y int) int {
	return x + y
}
