package main

import(
	"fmt"
)

func main() {
	fmt.Println(add(22, 43))
}

func add(x int, y int) int {
	return x + y
}
