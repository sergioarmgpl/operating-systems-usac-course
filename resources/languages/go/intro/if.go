package main //always start with package
import "fmt"

func main() {
	i := 0
	if i == 0 {
		fmt.Println("Zero")
	} else if i == 1 {
		fmt.Println("One")
	} else {
		fmt.Println("Other")
	}
}
