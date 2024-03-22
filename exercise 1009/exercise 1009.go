package main

import (
	"fmt"
)
 
func main() {
	var name string
    var	salFix, totVenda float64
	
	fmt.Scan(&name, &salFix, &totVenda)

	comit := totVenda * 15/100
	total := comit + salFix

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
