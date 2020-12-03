package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64 {
	var a float64
	a = 3.1416 * (math.Pow(radio, 2))
	return a
}

func numeroALetras(numero int) {
	//Switch
	var i int
	fmt.Println("Ingrese el número")
	fmt.Scanf("%d", &i)
	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	default:
		fmt.Println("Desconocido")
	}
}

//Múltiplos de 3 del 1 al 100
func multiplos3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	multiplos3()
	numeroALetras(2)
	fmt.Printf("El área del círculo con radio 3 es: %2.2f", areaCirculo(3))
}
