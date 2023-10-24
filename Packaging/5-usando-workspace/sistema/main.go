package main

import (
	"github.com/denisgariglio/goexpert/Packaging/5-usando-workspace/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
	println(uuid.New().String())
}
