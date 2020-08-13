package main

import "fmt"

func suma(a int, b int) int {
	return a + b
}

func holaMundo() {
	fmt.Println("Hola Mundo")
}

func leerTexto() {
	var texto string
	fmt.Println("Ingrese un texto: ")
	fmt.Scanf("%s", &texto)
	fmt.Printf("Texto = %s", texto)
}

func menu() {
	fmt.Println("Menú")
	fmt.Println("1. Sumar")
	fmt.Println("2. Restart")
	fmt.Println("3. Salir")
}

func main() {
	holaMundo()
	menu()
	leerTexto()
	resultado := suma(2, 3)
	fmt.Printf("\nLa suma es %d", resultado)
}
