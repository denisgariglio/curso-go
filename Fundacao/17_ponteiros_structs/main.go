package main

type Conta struct {
	saldo int
}

func (c *Conta) simular(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return c.saldo
}

// algo comum de utilizar ao criar uma struct que será utiliza por todo sistema
// crie apontando para a memoria e garantir que seja alterado na memoria e não na copia
func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func main() {
	conta := Conta{saldo: 100}
	conta.simular(200)
	println(conta.saldo)
}
