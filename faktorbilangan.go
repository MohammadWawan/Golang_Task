package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Printf("Masukkan angka = ")
	fmt.Scanln(&N)
	fmt.Println("Faktor dari", N, "adalah = ")
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			fmt.Println(i, "")
		}
	}
}
