package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	minhaVar1 := 10
	minhaVar2 := 20
	println(soma(&minhaVar1, &minhaVar2)) //nesse caso o & é a refernecia do valor na memoria
	println(minhaVar1)

}
