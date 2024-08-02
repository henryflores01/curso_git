package main

import (
	"fmt"
	"math"
)

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

func elevar(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func main() {
	var a, b int

	fmt.Println("Ingrese dos numeros: ")
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(sumar(a, b))
	fmt.Println(restar(a, b))
	fmt.Println(multiplicar(a, b))
	fmt.Println(dividir(a, b))
	fmt.Println(elevar(a, b))
}
