package main

import(
	"fmt"
)

const (
	Big = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFload(x float64) float64 {
	return x * 0.1
}
func main() {
	//fmt.Println("Big: %d, %08b", Big, Big)
	//fmt.Println("Small: %d, %08b", Small, Small)
	fmt.Println(needInt(Small))
	fmt.Println(needFload(Small))
	fmt.Println(needFload(Big))
}
