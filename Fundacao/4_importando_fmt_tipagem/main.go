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

	fmt.Printf("O tipo de E é %T", e) //%T tras o tipo da variavel

}
