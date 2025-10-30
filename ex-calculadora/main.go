package main

import "fmt"

type Calculadora struct {
	Operador1 float64
	Operador2 float64
}

func (c *Calculadora) Dividir() (float64,error) {
	if c.Operador2 == 0 {
		return 0, fmt.Errorf("Erro na divisão!")
	}
	result := c.Operador1 / c.Operador2
	return result, nil
}

func (c *Calculadora) Adicionar() (float64) {
	result := c.Operador1 + c.Operador2
	return result
}

func (c *Calculadora) Substrair() (float64) {
	result := c.Operador1 - c.Operador2
	return result
}

func (c *Calculadora) Multiplicar() (float64) {
	result := c.Operador1 * c.Operador2
	return result
}


func main()  {
	fmt.Println("Eis calculadora mestra!")
	var num1 float64
	var num2 float64
	var operacao string
	var result float64
	
	fmt.Printf("Digite o primeiro operador: ")
	fmt.Scanln(&num1)
	fmt.Printf("Digite a operação: ")
	fmt.Scanln(&operacao)
	fmt.Printf("Digite o segundo operador: ")
	fmt.Scanln(&num2)
			
	calculadora := Calculadora{num1, num2}
	
	switch operacao {
	case "+":
		result = calculadora.Adicionar()
	case "-":
		result = calculadora.Substrair()
	case "*":
		result = calculadora.Multiplicar()
	case "/":
		resultDivisao, err := calculadora.Dividir()
		if err != nil {
		fmt.Printf("Erro encontrado: %s", err)
		return
		}
		result = resultDivisao
	}

	fmt.Printf("O resultado da operação é : %f", result)


}
