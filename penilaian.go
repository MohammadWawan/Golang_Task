package main

import "fmt"

func main() {

	var nama_siswa string
	var score_siswa int

	fmt.Printf("Masukkan Nama :")
	fmt.Scanln(&nama_siswa)
	fmt.Printf("Masukkan Score:")
	fmt.Scanln(&score_siswa)
	fmt.Println("===============================")
	if score_siswa >= 80 {
		fmt.Println("Nama = ", nama_siswa)
		fmt.Printf("Score = A")
	} else if score_siswa >= 65 || score_siswa <= 79 {
		fmt.Println("Nama = ", nama_siswa)
		fmt.Printf("Score = B")
	} else if score_siswa >= 50 || score_siswa <= 64 {
		fmt.Println("Nama = ", nama_siswa)
		fmt.Printf("Score = C")
	} else if score_siswa >= 35 || score_siswa <= 49 {
		fmt.Println("Nama = ", nama_siswa)
		fmt.Printf("Score = D")
	} else if score_siswa >= 0 || score_siswa <= 34 {
		fmt.Println("Nama = ", nama_siswa)
		fmt.Printf("Score = E")
	} else {
		fmt.Println("Nama = %s", nama_siswa)
		fmt.Printf("Score = Nilai Invalid")
	}
}
