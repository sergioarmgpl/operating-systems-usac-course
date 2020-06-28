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
  go f(0)
  var input string
  fmt.Scanln(&input)
  fmt.Println(input)
}
