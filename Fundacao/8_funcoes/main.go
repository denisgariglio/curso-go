package main

import (
	"errors"
	"fmt"
)

func main() {

	// fmt.Println(sum(3, 7))
	// fmt.Println(sum1(43, 7))

	// fmt.Println(sumError(43, 7))

	// fmt.Println(sumError(3, 7))

	valor, err := sumError(50, 10)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}

}

func sum(a, b int) int {

	return a + b
}

// por conta de ser o mesmo tipo as variaveis, não precisa declarar mais de uma vez
func sum1(a, b int) (int, bool) {

	if a+b >= 50 {
		return a + b, true
	}

	return a + b, false
}

func sumError(a, b int) (int, error) {

	if a+b >= 50 {
		return 0, errors.New("O valor da soma é maior ou igual a 50")
	}

	return a + b, nil
}
