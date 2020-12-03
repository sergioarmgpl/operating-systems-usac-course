/*
1. Hacer función que devuelva 1 si los 2 números son iguales o 0 si no son iguales
2. Hacer la suma de los números de un arreglo de 10 elementos
3. Hacer una función que calcule el área de un círculo
4. Hacer un procedimiento que diga Bienvenido
5. Hacer un programa que muestre un mensaje según la siguiente tabla
- 11 Once
- 12 Doce
- 13 Trece
*/
package main //always start with package
import "fmt"

func sumaArreglo() {
	var total = 0
	var matriz = [10]int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	for _, valor := range matriz {
		total += valor
	}
	fmt.Printf("El total es %d\n", total)
}

func iguales(a int, b int) int {
	if a == b {
		return 1
	} else {
		return 0
	}

}

func area() {
	var h1, h2, g float64
	h1 = 0
	h2 = 3.14
	g = 0
	fmt.Printf("\n Ingrese el radio ")
	fmt.Scan(&h1)
	g = h1 * h1 * h2
	fmt.Printf("El área es: %2.1f \n", g)
}

func welcome() {
	fmt.Println("Bienvenido")
}

func mensaje(valor int) {
	//Switch
	switch valor {
	case 11:
		fmt.Println("Once")
	case 12:
		fmt.Println("Doce")
	case 13:
		fmt.Printf("Trece")
	}
}

func main() {
	var n1 int
	var n2 int
	fmt.Println("Ingresa n1:")
	fmt.Scan(&n1)
	fmt.Println("Ingresa n2:")
	fmt.Scan(&n2)
	fmt.Printf("%d\n", iguales(n1, n2))
	sumaArreglo()
	area()
	welcome()
	mensaje(12)
}
