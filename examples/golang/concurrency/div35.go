package main //always start with package

import "fmt"

func main() {
   var txt = ""
   for i:=1;i<=100;i++ {
      txt = ""
      if i % 3 == 0 { 
         txt += "Fizz"
      }
      if i % 5 == 0 {
         txt += "Buzz"
      }
      fmt.Println(i,txt)
   }
}
