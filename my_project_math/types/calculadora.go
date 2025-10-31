package types

import "errors"

type Calculadora struct {
	Operador1 float64
	Operador2 float64
}

func (c *Calculadora) Dividir() (float64, error) {
	if c.Operador2 == 0 {
		return 0, errors.New("erro na divisao: divisao por zero")
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
