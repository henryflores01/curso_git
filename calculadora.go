package main

import "fmt"

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

func main() {
	var a, b int

	fmt.Println("Ingrese dos numeros: ")
	fmt.Scanf("%d %d", &a, &b)

	fmt.Println(sumar(a, b))
	fmt.Println(restar(a, b))
	fmt.Println(multiplicar(a, b))
	fmt.Println(dividir(a, b))
}
