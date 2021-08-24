package main

import (
	"fmt"
)

func pangkat(base, pangkat int) int {
	output := 1
	for pangkat != 0 {
		output *= base
		pangkat -= 1
	}

	return output
}
func main() {
	fmt.Println(pangkat(7, 2))
}
