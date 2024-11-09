package main

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
	println(a)
	println(b)
	b = true
	println(b)
	println(c)
	println(d)
	println(e)

	//tipo de declar variável

	a := "Adriano Carvalho Batista"

	println(a)
	println(f)
}
