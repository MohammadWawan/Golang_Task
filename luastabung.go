package main

import "fmt"

func main() {

	const pi = 3.14
	var T, r float64
	fmt.Print("Masukkan jari-jari lingkaran =")
	fmt.Scanln(&r)
	fmt.Print("Masukkan tinggi tabung =")
	fmt.Scanln(&T)

	fmt.Print("Luas tabung =", 2*pi*r*(r+T))

	/*
		const pi = 3.14

		fmt.Println("Masukkan Tinggi dan Jari-jari Alas Tabung:")
		fmt.Println("*pisahkan dua nilai dengan spasi")
		var T, r float64
		fmt.Scanf("%v", &T)
		fmt.Scanf("%v", &r)

		L := 2 * pi * r * (r + T)
		fmt.Println(L)
	*/
}
