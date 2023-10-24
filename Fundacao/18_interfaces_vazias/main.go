package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, Wordl!"

	ShowType(x)
	ShowType(y)
}

func ShowType(t interface{}) {
	fmt.Printf("O tipo da variavel é %T e o valor é %v\n", t, t)
}
