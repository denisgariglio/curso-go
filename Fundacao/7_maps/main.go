package main

import "fmt"

func main() {

	salarios := map[string]int{"Denis": 1000, "João": 2000, "Jose": 3000}
	// fmt.Println(salarios["Denis"])
	// delete(salarios, "Denis")
	// fmt.Println(salarios["Denis"])
	// salarios["Gariglio"] = 5000
	// fmt.Println(salarios["Gariglio"])

	// sal := make(map[string]int)
	// sal1 := map[string]int{}
	// sal["Gariglio"] = 5000
	// sal["Gar"] = 5001
	// fmt.Println(sal)
	// sal1["Gariglio"] = 5000
	// fmt.Println(sal1)

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios { //_ é conhecido como blank identifier
		fmt.Printf("O salario é %d\n", salario)
	}

}
