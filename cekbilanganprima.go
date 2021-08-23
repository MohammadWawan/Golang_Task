package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Println("Masukkan Bilangan :")
	fmt.Scanln(&bilangan)

	if bilangan <= 1 {
		fmt.Println("Masukkan Angka lebih dari 1.")
	}
	if bilangan == 2 {
		fmt.Println(bilangan, "adalah bilangan prima")
	} else {
		for i := 2; i < bilangan; i++ {
			if bilangan%i == 0 {
				fmt.Println(bilangan, "adalah bukan bilangan prima")
				break
			} else {
				fmt.Println(bilangan, "adalah bilangan prima")
			}
		}
	}

}
