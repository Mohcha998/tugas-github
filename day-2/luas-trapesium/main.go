package main

import "fmt"

func rumusTrapesium(atas, bawah, tinggi float64) float64 {
	return 0.5 * (atas + bawah) * tinggi
}

func main() {
	atas := 5
	bawah := 10
	tinggi := 6

	luas := rumusTrapesium(float64(atas), float64(bawah), float64(tinggi))
	fmt.Println("Luas Trapesium adalah :", luas)
}
