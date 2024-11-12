package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	adriano := Cliente{
		Nome:  "Adriano Carvalho Batista",
		Idade: 37,
		Ativo: true,
	}

	fmt.Println(adriano.Nome)
}
