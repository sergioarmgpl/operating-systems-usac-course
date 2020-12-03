package main //always start with package
import "fmt"

func main() {
	var n1, n2 int16
	n1 = 0
	n2 = 0
	fmt.Printf("Ingrese n1: \n")
	fmt.Scanf("%d\n", &n1)

	fmt.Printf("Ingrese n2: \n")
	fmt.Scanf("%d\n", &n2)

	fmt.Printf("\n %d + %d = %d", n1, n2, n1+n2)
}
