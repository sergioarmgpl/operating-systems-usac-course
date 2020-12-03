package main

import (
	"fmt"
)

func main() {
	var numero int16
	fmt.Printf("Ingrese el número: \n")
	fmt.Scanf("%d", &numero)
	//Comparar el valor
	if numero < 25 {
		fmt.Println("Es menor 25")
	}
	//If-Else Igual
	if numero%2 == 0 {
		fmt.Println("Es par")
	} else {
		fmt.Println("Es impar")
	}
	// número not equal 25
	if numero != 25 {
		fmt.Println("No es 25")
	} else {
		fmt.Println("Es 25")
	}
	//If-If-Else-If-Else-Else anidado(Nested)
	if numero != 100 {
		fmt.Println("No es igual 100")
	} else if numero != 30 {
		fmt.Println("No es 30")
	} else if numero >= 300 {
		fmt.Println("Es mayor a 300")
	}
}
