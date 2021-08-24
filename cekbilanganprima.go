package main

import "fmt"

func primeNumber(number int) bool {
	for x := 2; x > 1; x++ {
		if number%2 == 0 {
			return false
		} else {
			return true
		}
	}
	return true
}

func main() {
	fmt.Printf("%v", primeNumber(4))
}
