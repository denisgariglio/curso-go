package main

func main() {

	//Memória -> Endereço -> Valor

	a := 10

	println(&a)

	// o * é referente ao valor do endereço da memoria
	var ponteiro *int = &a
	println(ponteiro)

	*ponteiro = 20

	b := &a

	println(*b)

	*b = 30
	println(a)
}
