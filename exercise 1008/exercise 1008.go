package main

import (
	"fmt"
)
 
func main() {
    var	salHour float32
	var numFunc, numHour int
	
	fmt.Scan(&numFunc, &numHour, &salHour)

	salary := salHour * float32(numHour)
	fmt.Println("NUMBER =", numFunc)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
