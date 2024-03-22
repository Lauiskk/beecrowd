package main

import (
	"fmt"
)
 
func main() {
	var (
		codigP1, quantP1, codigP2, quantP2 int
		valP1, valP2 float32
	)

	fmt.Scan(&codigP1, &quantP1, &valP1)
	fmt.Scan(&codigP2, &quantP2, &valP2)

	valTot := (valP1 * float32(quantP1))+ (valP2 * float32(quantP2))

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", valTot)


}
