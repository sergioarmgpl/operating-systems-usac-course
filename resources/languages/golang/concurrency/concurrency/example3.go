package main

import (
	"fmt"
	"time"
)

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
		time.Sleep(time.Millisecond * 1000)
	}
}

func f2() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	go f1()
	go f2()
	var input string
	fmt.Println("Enter a word")
	fmt.Scanln(&input)
	fmt.Println(input[0:3])
}
