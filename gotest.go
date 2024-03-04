package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func main() {
	angka1 := 5
	angka2 := 3
	hasil := tambah(angka1, angka2)

	fmt.Printf("%d + %d = %d\n", angka1, angka2, hasil)
}
