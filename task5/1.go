package main

import (
	"fmt"
)

func main() {
	fmt.Println(primeX(1)) // 2
}
func primeX(number int) int {
	i := 2
	x := 0
	for {
		a := 0
		for j := 2; j < i; j++ {
			a++
		}
		if a == 0 && x < number {
			x++
		}
		if x == number {
			break
		}
		i++
	}
	return i
}
