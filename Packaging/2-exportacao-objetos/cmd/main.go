package main

import (
	"fmt"

	"github.com/denisgariglio/goexpert/Packagin/2-exportacao-objetos/math"
)

func main() {
	m := math.NewMath(1, 3)
	fmt.Println(m.Add())
}
