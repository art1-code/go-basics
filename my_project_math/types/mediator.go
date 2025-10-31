package types

import (
	"errors"
	"fmt"
)

type Mediator struct {
	calculadora *Calculadora
	operacao    string
}

// cria função para receber valores
func (m *Mediator) SolicitarValores() error {
	fmt.Println("Digite o primeiro número:")
	if _, err := fmt.Scanln(&m.calculadora.Operador1); err != nil {
		return errors.New("erro na primeira entrada")
	}
	fmt.Println("Digite o segundo número:")
	if _, err := fmt.Scanln(&m.calculadora.Operador2); err != nil {
		return errors.New("erro na segunda entrada")
	}
	return nil
}

// cria funcao para receber operacao desejada
func (m *Mediator) SolicitarOperacao() error {
	fmt.Println("Digite a operacao desejada:")
	var operacao string
	if _, err := fmt.Scanln(&operacao); err != nil {
		return errors.New("erro na entrada da operação")
	}

	switch operacao {
	case "+", "-", "/", "*":
		m.operacao = operacao
	default:
		return errors.New("erro na operacao")
	}
	return nil
}

// realiza a operacao de acordo com os parametros passados
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
		} else {
			fmt.Println(err)
		}
	default:
		break
	}

	return nil
}

// executa todas as funcs
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

// funcao usada para criar um novo Mediator
func NewMediator(c *Calculadora) *Mediator {
	return &Mediator{calculadora: c}
}
