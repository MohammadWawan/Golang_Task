package main

import (
	"fmt"
	"math"
)

func main() {
	var N float64
	fmt.Println("Masukkan Bilangan :")
	fmt.Scanln(&N)

	var s = math.Round(math.Sqrt(N))

	if s*s != N {
		fmt.Println("Lampu Mati /n")
	} else {
		fmt.Println("Lampu Menyala /n")
	}
}
