package main

import "fmt"

func main() {
	angka := 4
	nilai := 0

	for i := 1; i <= angka; i++ {
		if angka%i == 0 {
			nilai++
		}
	}
	if nilai == 2 {
		fmt.Println(angka, "Adalah Bilangan Prima")
	} else {
		fmt.Println(angka, "Bukan Bilangan Prima")
	}
}
