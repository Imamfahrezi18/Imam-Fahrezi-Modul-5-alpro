package main

import "fmt"

func printbintang(n int) {
	if n > 0 {
		fmt.Print("*")
		printbintang(n - 1)
	}
}

func printPattern(n, current int) {
	if current <= n {
		printbintang(current)
		fmt.Println()
		printPattern(n, current+1)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah baris: ")
	fmt.Scan(&n)

	fmt.Println("Pola bintang:")
	printPattern(n, 1)
}
