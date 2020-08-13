package main

import "fmt"

func main() {
	i := 4
	//If
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("Other")
	}

	//Switch
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2, 4:
		fmt.Printf("%d is Even", i)
	default:
		fmt.Println("Other")
	}
}
