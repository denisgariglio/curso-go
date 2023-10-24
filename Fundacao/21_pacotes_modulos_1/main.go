package main

//go mod init --> inicia um gerenciador de pacotes
import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {

	carro := matematica.Carro{Marca: "Toyota"}
	fmt.Println(carro.Marca)
	fmt.Println(carro.Andar())
	s := matematica.Soma(10, 20)
	fmt.Println("Resultado ", s)
	fmt.Println(matematica.A)
	fmt.Println(uuid.New())
}
