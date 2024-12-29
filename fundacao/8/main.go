package main

import (
	"errors"
	"fmt"
)

func main() {

	fmt.Println(sum(10, 10))
	fmt.Println(sum_two_return(25, 25))

	valor, err := sum_two_return_erros(10, 90)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)

	}
}

func sum(a int, b int) int {
	return a + b
}

func sum_two_return(a int, b int) (int, bool) {
	if a+b >= 50 {
		return a + b, true
	}
	return a + b, false
}

func sum_two_return_erros(a int, b int) (int, error) {
	if a+b >= 50 {
		return a + b, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}
