package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {

	denis := Cliente{
		Nome:  "Denis",
		Idade: 40,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, Idade: %d, Ativo: %t\nr", denis.Nome, denis.Idade, denis.Ativo)

}
