package main

import "fmt"

func main()  {
	
	var numero1, numero2 int

	fmt.Println("Digite o primeiro numeral da soma:")
	fmt.Scanln(&numero1)
	fmt.Println("Digite o segundo numeral da soma:")
	fmt.Scanln(&numero2)

	soma := func(num1, num2 int) int {
		return num1 + num2
	}

	sum := soma(numero1, numero2)

	fmt.Printf("O soma de %d e %d é: %d\n", numero1, numero2, sum)
}