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

func main() {

	denis := Cliente{
		Nome:  "Denis",
		Idade: 40,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\nr", denis.Nome, denis.Idade, denis.Ativo)

}
