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
        fmt.Println("Proceso 0: turno:",turno)
        fmt.Println("Proceso 0: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    time.Sleep(time.Millisecond * 1000)
    fmt.Println("Proceso 0: Sección Crítica")
    turno = 1
    fin <- true
}

func proceso1 () {
    for turno != 1 {
        fmt.Println("Proceso 1: turno:",turno)
        fmt.Println("Proceso 1: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    time.Sleep(time.Millisecond * 1000)
    fmt.Println("Proceso 1: Sección Crítica")
    turno = 0
    fin <- true
}

func main () {
   for {
      go proceso0()
      go proceso1()
      time.Sleep(time.Millisecond * 1000)
   }
   <- fin
}
