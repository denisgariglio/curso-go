package main

import "github.com/denisgariglio/goexpert/Packaging/4-gomod-replace/math"

func main() {
	m := math.Math{A: 1, B: 2}
	println(m.Add())
}
