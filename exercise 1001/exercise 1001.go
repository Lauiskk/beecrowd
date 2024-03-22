package main

import (
	"fmt"
)

func add (num1 int, num2 int) int{
	return num1 + num2
}

func main(){
	var numS1, numS2 int

	fmt.Scan(&numS1, &numS2)

	fmt.Println(add(numS1, numS2))
}