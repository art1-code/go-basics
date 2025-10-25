package main

import "fmt"

func main()  {
	
	var age int

	fmt.Println("Digite sua idade:")
	fmt.Scanln(&age)

	if age > 18 {
		fmt.Printf("Vc pode ser preso, meu chapa")
	} else {
		fmt.Printf("vc vai pra o lar do garoto e jaja ta solto")
	}
}