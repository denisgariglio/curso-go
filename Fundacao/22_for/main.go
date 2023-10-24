package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	// numeros := []string{"um", "dois", "três"}
	// for k, v := range numeros {
	// 	println(k, v)
	// }

	numeros := []string{"um", "dois", "três"}
	for _, v := range numeros {
		println(v)
	}

	numeros2 := []string{"um", "dois", "três"}
	for k, _ := range numeros2 {
		println(k)
	}

	i := 0

	for i < 10 {
		println(i)
		i++
	}

	// for {
	// 	println("Hello World")
	// }

}
