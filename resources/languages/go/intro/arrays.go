package main

import "fmt"

func main() {
	//arreglo = [1,6,7,82]
	//for i in range(0,4):
	//   print(arreglo[i])
	//for dato in arreglo:
	//   print(dato)
	var x [3]int
	x[0] = 3
	x[1] = 33
	x[2] = 333
	//[0 0 0 0 100]
	fmt.Println(x[2])
	fmt.Println("--------------------")
	//Other form of declaration
	var total = 0
	var y = [5]int{98, 93, 77, 82, 83}
	for _, value := range y {
		fmt.Println(value)
		total += value
	}
	fmt.Printf("total = %d\n", total)
	fmt.Println("--------------------")
	var L int
	L = len(y)
	for i := 0; i <= (L - 1); i++ {
		fmt.Println(y[i])
	}
}
