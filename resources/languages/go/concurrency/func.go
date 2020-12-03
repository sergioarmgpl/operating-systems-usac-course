package main

import "fmt"

func add(a int,b int) int {
   return a+b
}

func proc1() {
  fmt.Println("I m here")
}

func main() {
  proc1()
  fmt.Println("add: ",add(3,3))
}

