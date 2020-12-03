package main

import "fmt"

func main() {
   var x [3]int
   x[0]=1
   x[1]=3
   x[2]=3
   for _,value := range x {
      fmt.Println(value)
      //total += value
   }
   y := make([]int, 5, 9)
   y[0]=2
   y[1]=4
   y[2]=9
   y[3]=7
   y[4]=11
   for _,value := range y {
      fmt.Println(value)
      //total += value
   }
   fmt.Println(y[1:3])
}

