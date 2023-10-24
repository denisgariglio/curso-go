package main

import "fmt"

const a = "Hello World !!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Denis"
	e float64 = 1.2
	f ID      = 3
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30

	fmt.Println(len(meuArray) - 1)
	fmt.Println(len(meuArray))
	fmt.Println(meuArray)
	fmt.Println(meuArray[0])
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor do indice é %d e o valor é %d\n", i, v)
	}

}
