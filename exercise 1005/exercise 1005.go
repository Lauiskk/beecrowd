package main

import "fmt"

func main() {
	var A, B float32

	fmt.Scan(&A, &B)

	var MEDIA = ((A * 3.5)+(B * 7.5))/11
	
	fmt.Printf("MEDIA = %.5f\n", MEDIA)
}