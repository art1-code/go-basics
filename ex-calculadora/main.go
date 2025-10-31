package main

import (
	"errors"
	"fmt"
)

type Calculadora struct {
	Operador1 float64
	Operador2 float64
}

type Runner struct {
	calculadora *Calculadora
	resultado   float64
	operacao    string
}

func (c *Calculadora) Dividir() (float64, error) {
	if c.Operador2 == 0 {
		return 0, errors.New("Erro na divisão!")
	}
	return (c.Operador1 / c.Operador2), nil
}

func (c *Calculadora) Adicionar() float64 {
	return c.Operador1 + c.Operador2
}

func (c *Calculadora) Substrair() float64 {
	return c.Operador1 - c.Operador2
}

func (c *Calculadora) Multiplicar() float64 {
	return c.Operador1 * c.Operador2

}

func (r *Runner) SolicitaOperadores() error {
	fmt.Println("Digite o primeiro operador:")
	_, err1 := fmt.Scanln(&r.calculadora.Operador1)
	if err1 != nil {
		return errors.New("Operador não é válido")
	}

	fmt.Println("Digite o segundo operador:")
	_, err2 := fmt.Scanln(&r.calculadora.Operador2)
	if err2 != nil {
		return errors.New("Operador não é válido")
	}

	return nil
}

func (r *Runner) SolicitaOperacao() (float64, error) {

	fmt.Println("Digite a operacao:")
	_, err := fmt.Scanln(&r.operacao)
	if err != nil {
		return 0, errors.New("Erro encontrado durante a entrada")
	}
	switch r.operacao {
	case "+", "/", "-", "*":
		return r.ExecutaOperacao()
	default:
		return 0, errors.New("Operacao Inválida")
	}
}

func (r *Runner) ExecutaOperacao() (float64, error) {
	switch r.operacao {
	case "+":
		r.resultado = r.calculadora.Adicionar()
	case "-":
		r.resultado = r.calculadora.Substrair()
	case "*":
		r.resultado = r.calculadora.Multiplicar()
	case "/":
		result, err := r.calculadora.Dividir()
		if err != nil {
			return 0, errors.New("Erro encontrado")
		}
		r.resultado = result
	}
	return r.resultado, nil
}

func (r *Runner) Execute() {
	r.SolicitaOperadores()

	r.SolicitaOperacao()
}

func NewRunner(c *Calculadora) *Runner {
	return &Runner{c, 0, ""}
}

func main() {
	fmt.Println("Eis calculadora mestra!")
	calculadora := &Calculadora{}
	r := NewRunner(calculadora)
	r.Execute()
	fmt.Printf("O resultado da operação é : %f", r.resultado)

}
