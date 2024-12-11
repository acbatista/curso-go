package main

import "fmt"

type Cliente struct {
	nome string
}

func (c Cliente) andou() {
	c.nome = "Adriano Carvalho Batista"
	fmt.Printf("%s o cliente andou.\n", c.nome)
}

func main() {
	adriano := Cliente{"adriano"}
	adriano.andou()

	fmt.Printf("%s O valor da struct\n", adriano.nome)
}
