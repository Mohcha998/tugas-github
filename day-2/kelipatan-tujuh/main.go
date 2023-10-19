package main

import "fmt"

func Search(angka int) bool {
	if angka%7 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		if Search(i) {
			fmt.Println(i, "Adalah Kelipatan 7")
		}
	}

}
