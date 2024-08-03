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

	fmt.Printf("%d + %d = %d\n", a, b, sumar(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, restar(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, multiplicar(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, dividir(a, b))
}
