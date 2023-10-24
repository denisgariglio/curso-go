package main

type MyNumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T Number](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Denis": 1000, "Joao": 2000, "Jose": 3000}
	println(SomaInteiro(m))

	m2 := map[string]float64{"Denis": 1000.45, "Joao": 200.0, "Jose": 3000.15}
	println(SomaFloat(m2))

	println(Soma(m))
	println(Soma(m2))

	m3 := map[string]MyNumber{"Denis": 12000, "Joao": 2000, "Jose": 3000}
	println(Soma(m3))
	println(Compara(10, 12))

}
