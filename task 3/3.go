package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {
	hasil := make(map[string]int)
	a := append(arrayA, arrayB...)
	res := make([]string, 0)
	for _, val := range a {
		hasil[val] = 1
	}
	for letter, _ := range hasil {
		res = append(res, letter)
	}
	return res
}

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	// [“king”, “devil jin”, “akuma”, “eddie”, “steve”, “geese”]
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	// [“sergei”, “jin”, “steve”, “bryan”]

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	// [“alisa”, “yoshimitsu”, “devil jin”, “law”]

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	// [“devil jin”, “sergei”]

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	// [“hwoarang”]

	fmt.Println(ArrayMerge([]string{}, []string{}))

	// []
}
