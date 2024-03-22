package main

import "fmt"

func main() {
	var A, B, C, D int

	fmt.Scan(&A, &B, &C, &D)

	DIFERENCA := (A * B - C * D)
	
	fmt.Println("DIFERENCA =", DIFERENCA)
}