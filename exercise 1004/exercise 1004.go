package main

import "fmt"

func main() {
	var A, B int

	fmt.Scan(&A, &B)

	PROD := A * B
	
	fmt.Println("PROD =", PROD)
}