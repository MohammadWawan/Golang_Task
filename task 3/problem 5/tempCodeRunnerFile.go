package main

import (
	"fmt"

	mapset "github.com/deckarep/golang-set"
)

func PairSum(arr []int, target int) []int {
	indexValue := map[int]int{}
	for i, v := range arr {
		indexValue[v] = i
	}

	set := mapset.NewSet()
	for i := 0; i < len(arr); i++ {
		temp := target - arr[i]
		if set.Contains(temp) {
			indexOfTemp := indexValue[temp]
			return []int{indexOfTemp, i}
		}
		set.Add(arr[i])
	}
	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
