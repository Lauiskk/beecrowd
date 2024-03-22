package main

import (
	"fmt"
)
 
func main() {
	var N float64

	fmt.Scan(&N)

	var nota100, nota50, nota20, nota10, nota5, nota2, moeda1,moeda50,moeda25,moeda10,moeda5,moeda01 int

	for (N >= 0 && N <= 1000000.00) {
		if N >= 100 {
			nota100++
			N -= 100
		}else if N >= 50 && N < 100{
			nota50++
			N -= 50
		}else if N >= 20 && N < 50{
			nota20++
			N -= 20
		}else if N >= 10 && N < 20{
			nota10++
			N -= 10
		}else if N >= 5 && N < 10{
			nota5++
			N -= 5
		}else if N >= 2 && N < 5{
			nota2++
			N -= 2
		}else if N >= 1 && N < 2{
			moeda1++
			N -= 1
		}else if N >= 0.50 && N < 1{
			moeda50++
			N -= 0.50
		}else if N >= 0.25 && N < 0.50{
			moeda25++
			N -= 0.25
		}else if N >= 0.10 && N < 0.25{
			moeda10++
			N -= 0.10
		}else if N >= 0.05 && N < 0.10{
			moeda5++
			N -= 0.05
		}else if N >= 0.009 && N < 0.05{
			moeda01++
			N -= 0.01
		}else{
			break
		}
		
	}
	fmt.Println("NOTAS:")
	fmt.Println(nota100, "nota(s) de R$ 100.00")
	fmt.Println(nota50, "nota(s) de R$ 50.00")
	fmt.Println(nota20, "nota(s) de R$ 20.00")
	fmt.Println(nota10, "nota(s) de R$ 10.00")
	fmt.Println(nota5, "nota(s) de R$ 5.00")
	fmt.Println(nota2, "nota(s) de R$ 2.00")
	fmt.Println("MOEDAS:")
	fmt.Println(moeda1, "moeda(s) de R$ 1.00")
	fmt.Println(moeda50, "moeda(s) de R$ 0.50")
	fmt.Println(moeda25, "moeda(s) de R$ 0.25")
	fmt.Println(moeda10, "moeda(s) de R$ 0.10")
	fmt.Println(moeda5, "moeda(s) de R$ 0.05")
	fmt.Println(moeda01, "moeda(s) de R$ 0.01")
}