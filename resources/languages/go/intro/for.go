package main //always start with package

import "fmt"

func main() {
	//Example1
	//i+=1 ==> i++ ==> i = i + 1
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
	//For --> While
	//Example2
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
