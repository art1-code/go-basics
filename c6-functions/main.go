package main

import "fmt"

func somarPositivos(num1, num2 int) (int, error) {

	if num1 < 0 || num2 < 0 {
		return 0, fmt.Errorf("Esta função somente soma valores positivos")
		
	}	
	soma := num1 + num2
	return soma,nil
}

func main()  {

	var numero1 int
	var numero2 int
	var attempts int = 3

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

		soma, err := somarPositivos(numero1,numero2)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("O valor da soma é: %d\n", soma)
		}

		break
	}

}