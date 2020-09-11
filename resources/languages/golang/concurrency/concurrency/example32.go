package main

import (
	"fmt"
	"time"
)

var done = make(chan bool)

func f1() {
	for i := 0; i < 4; i++ {
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
	done <- true
}

func main() {
	go f1()
	go f2()
	<-done
}
