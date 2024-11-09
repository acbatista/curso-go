package main

import "fmt"

func main() {
	salarios := map[string]int{"Adriano": 100, "Adryan": 200, "Amanda": 300}

	fmt.Println(salarios["Adryan"])

	// deletando uma posição do map
	delete(salarios, "Adriano")
	salarios["Nego"] = 1000

	fmt.Println(salarios["Nego"])

	//função make, para preparar map

	sal := make(map[string]int)
	sal["sal"] = 10

	fmt.Println(sal["sal"])

	for nome, salario := range salarios {
		fmt.Printf("O salario de %s é %d \n", nome, salario)
	}

	// ignore black indent file

	for _, salario := range salarios {
		fmt.Printf("O salario é %d \n", salario)
	}

}
