package main

import "fmt"

func main() {
	var minhaVar interface{} = "Adriano Carvalho"

	//converter
	println(minhaVar.(string))

	res, ok := minhaVar.(int)

	fmt.Printf("Valor de res: %v valor de ok é igual %v \n", res, ok)

	res2 := minhaVar.(int)

	//retorna um erro panic
	fmt.Printf("O valor de res2 é igual %v \n", res2)
}
