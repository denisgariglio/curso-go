package main

import (
	"fmt"

	"github.com/denisgariglio/goexpert/Packagin/1-go-mod/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
