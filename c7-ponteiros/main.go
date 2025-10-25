package main

import "fmt"

func main()  {

	var numero1 int
	var numero2 int
	var attempts int = 3
	var soma int

	for {
		
		if attempts == 0 {
			fmt.Println("Número de tentativas excedido!")
			break
		}
		fmt.Printf("Digite o primeiro número:")
		_, err1 := fmt.Scanf("%d\n", &numero1)
		fmt.Printf("Digite o segundo número:")
		_, err2 := fmt.Scanf("%d\n",&numero2)

		if err1 != nil || err2 != nil {
			fmt.Println("Uma (ou mais) entrada(s) inválida(s). Tente novamente")
			attempts-- 
			fmt.Printf("Número de tentativas restantes: %d\n", attempts)
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		
		somarPositivos := func (soma *int, num1 *int, num2 *int) {
			if *num1 < 0 || *num2 < 0 {
				fmt.Println("Valores negativos não serão tolerados!!!")
				*soma = 0
				return
			}
			*soma = *num1 + *num2
		}

		somarPositivos(&soma, &numero1, &numero2)
		
		if soma == 0 {
			break
		}

		fmt.Printf("O valor da soma é: %d\n", soma)

		break
	}

}