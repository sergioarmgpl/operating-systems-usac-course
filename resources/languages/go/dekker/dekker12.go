package main
/* producer-consumer problem in Go */


import (
   "fmt"
   "time"
)

var fin = make(chan bool)
var turno = 0

func proceso0 () {
    for turno != 0 {
        fmt.Println("Proceso 0: turno: ", turno)
        fmt.Println("Proceso 0: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    fmt.Println("Proceso 0: turno: ", turno)
    fmt.Println("Proceso 0: Sección Crítica")
    time.Sleep(time.Millisecond * 1000)
    turno = 1
    fin <- true // Signal completion
}

func proceso1 () {
    for turno != 1 {
        fmt.Println("Proceso 1: turno: ", turno)
        fmt.Println("Proceso 1: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    fmt.Println("Proceso 1: turno: ", turno)
    fmt.Println("Proceso 1: Sección Crítica")
    time.Sleep(time.Millisecond * 1000)
    turno = 0
    fin <- true // Signal completion
}

func main () {
   for { // Run 5 iterations as example
      go proceso0()
      <-fin // Wait for proceso0 to finish
      go proceso1()
      <-fin // Wait for proceso1 to finish
   }
}
