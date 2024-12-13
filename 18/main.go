package main

import "fmt"

//func Soma(m map[string]int) int {
//	var soma int
//	for _, v := range m {
//		soma += v
//	}
//
//	return soma
//
//}

//func Soma[T int | float64](m map[string]T) T {
//	var soma T
//	for _, v := range m {
//		soma += v
//	}
//
//	return soma
//
//}

//usando constraint

type MyNumber int

type Number interface {
	~int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma

}

func main() {

	m := map[string]int{"teste": 100, "test2": 200, "test3": 300}
	m2 := map[string]float64{"teste": 100.03, "test2": 200.3, "test3": 300.4}

	m3 := map[string]MyNumber{"teste": 100, "test2": 200, "test3": 300}

	fmt.Println(Soma(m))
	fmt.Println(Soma(m2))
	fmt.Println(Soma(m3))

}
