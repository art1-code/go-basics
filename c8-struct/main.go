package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade int
	Email string
	Senha string
}

func (pessoa *Usuario) Saudacao() string {
	saudacao := "Olá, ma chamo " + pessoa.Nome
	return saudacao
}

func main() {

	pessoa := Usuario{"Arthur", 27, "arthur1silva1r@gmail.com", "teste123"}
	fmt.Println(pessoa.Nome)
	fmt.Println(pessoa.Idade)
	fmt.Println(pessoa.Email)
	fmt.Println(pessoa.Senha)
	fmt.Println(pessoa.Saudacao())

}
