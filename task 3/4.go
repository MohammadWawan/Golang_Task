package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var b = make(map[string]string)
	var stringNumValue string

	//input angka pertama ke map
	stringNumValue = string(angka[0])
	b[stringNumValue] = stringNumValue

	//perulangan setiap nomor dalam string
	for i := 1; i < len(angka); i++ {
		stringNumValue = string(angka[i])
		if _, ok := b[stringNumValue]; ok {
			delete(b, stringNumValue)
		} else {
			b[stringNumValue] = stringNumValue
		}
	}

	//konversi map ke slice
	var result []int
	for key, _ := range b {
		keyInt, _ := strconv.Atoi(key)
		result = append(result, keyInt)
	}
	return result
}

func main() {
	fmt.Println(munculSekali("1234123")) // [4]

	fmt.Println(munculSekali("76523752")) // [6, 3]

	fmt.Println(munculSekali("12345")) // [1 2 3 4 5]

	fmt.Println(munculSekali("1122334455")) // []

	fmt.Println(munculSekali("0872504")) // [8 7 2 5 4]
}
