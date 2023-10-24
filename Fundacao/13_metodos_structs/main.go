package main

import "fmt"

type Endereco struct {
	Numero int
	Cidade string
	Estado string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {

	denis := Cliente{
		Nome:  "Denis",
		Idade: 40,
		Ativo: true,
	}

	denis.Desativar()

	// fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\nr", denis.Nome, denis.Idade, denis.Ativo)

}
