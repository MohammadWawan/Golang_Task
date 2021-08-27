package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	if number <= 1 {
		return false
	}
	sq_root := int(math.Sqrt(float64(number)))
	for i := 2; i < sq_root; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
func isprime(number int) string {
	if primeNumber(number) == true {
		fmt.Print("bilangan prima")
	} else {
		fmt.Print("bukan bilangan prima")
	}
	return ""
}
func main() {
	fmt.Println(isprime(1000000007)) // true

	fmt.Println(isprime(1500450271)) // true

	fmt.Println(isprime(1000000000)) // false

	fmt.Println(isprime(10000000019)) // true

	fmt.Println(isprime(10000000033)) // true

}
