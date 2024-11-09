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
	fmt.Printf("O tipo de F é %T ", f) // %T tipo
	fmt.Printf("O tipo de F é %v ", f) // %v valor
}
