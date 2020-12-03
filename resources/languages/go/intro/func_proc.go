package main

import "fmt"

func holaMundo() {
	fmt.Println("Hola Mundo")
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
}
