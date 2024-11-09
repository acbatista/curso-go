package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Adriano"
	e float64 = 1.2
	f ID      = 100
)

func main() {
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Printf("O tipo de F é %v ", f) // %v valor
	fmt.Println()
	fmt.Println(meuArray[1])

	//pecorrendo o array

	for i, v := range meuArray {
		fmt.Printf("O valor de indice é %d e o valor é %d \n", i, v)
	}
}
