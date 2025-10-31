package main

import (
	"errors"
	"fmt"
)

// ----------------- Calculadora ----------------

type Calculadora struct {
	Operador1 float64
	Operador2 float64
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

// ----------------- Mediator ----------------

// Cria struct que vai mediar as ações na calculadora
type Mediator struct {
	calculadora *Calculadora
	operacao    string
}

// cria função para receber valores
func (m *Mediator) SolicitarValores() error {
	fmt.Println("Digite o primeiro número:")
	if _, err := fmt.Scanln(&m.calculadora.Operador1); err != nil {
		return errors.New("Erro na primeira entrada")
	}
	fmt.Println("Digite o segundo número:")
	if _, err := fmt.Scanln(&m.calculadora.Operador2); err != nil {
		return errors.New("Erro na segunda entrada")
	}
	return nil
}

// cria funcao para receber operacao desejada
func (m *Mediator) SolicitarOperacao() error {
	fmt.Println("Digite a operacao desejada:")
	var operacao string
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("Erro na entrada da operacao")
	}

	switch operacao {
	case "+", "-", "/", "*":
		m.operacao = operacao
	default:
		return errors.New("Operacao inválida")
	}
	return nil
}

func (m *Mediator) ExacutaOperacao() error {
	switch m.operacao {
	case "+":
		fmt.Println("Resultado: ", m.calculadora.Adicionar())
	case "-":
		fmt.Println("Resultado: ", m.calculadora.Substrair())
	case "*":
		fmt.Println("Resultado: ", m.calculadora.Multiplicar())
	case "/":
		if resultado, err := m.calculadora.Dividir(); err == nil {
			fmt.Println("Resultado: ", resultado)
		}
	default:
		break
	}

	return nil
}

func (m *Mediator) Execute() {
	for {
		if err := m.SolicitarValores(); err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		if err := m.SolicitarOperacao(); err != nil {
			fmt.Println("Erro:", err)
			continue
		}

		if err := m.ExacutaOperacao(); err != nil {
			fmt.Println("Erro:", err)
			continue
		}
	}
}

func NewMediator(c *Calculadora) *Mediator {
	return &Mediator{calculadora: c}
}

// ----------------- Função Main ----------------
func main() {
	calculadora := Calculadora{}
	mediator := NewMediator(&calculadora)
	mediator.Execute()
}
