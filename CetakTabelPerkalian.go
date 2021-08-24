package main

import (
	"fmt"
)

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Printf("%d", i*j)

		}
		fmt.Printf("\n")
	}
	return

}
func main() {
	cetakTablePerkalian(9)
}
