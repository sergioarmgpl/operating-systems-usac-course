package main //always start with package

import "fmt"

func main() {
   for i:=1;i<=100;i++ {
      if i % 3 == 0 {
         fmt.Println(i)
      }
   }
}
