package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     string
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	//Endereco
	Addres Endereco
}

func main() {
	adriano := Cliente{
		Nome:  "Adriano Carvalho Batista",
		Idade: 37,
		Ativo: true,
	}

	adriano.Addres.Estado = "DF"
	adriano.Addres.Cidade = "Santa Maria"
	adriano.Addres.Logradouro = "Total Ville"
	adriano.Addres.Numero = "203"

	fmt.Println(adriano.Nome)
	fmt.Println(adriano.Addres.Estado)
}
