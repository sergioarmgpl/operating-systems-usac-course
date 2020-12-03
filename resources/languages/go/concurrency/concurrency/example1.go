package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	var input string
	go f(0)
	fmt.Println("Input a number:")
	fmt.Scanln(&input)
	fmt.Println(input)
}
