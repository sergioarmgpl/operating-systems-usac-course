package main
/* producer-consumer problem in Go */


import (
   "fmt"
   "time"
)

var fin = make(chan bool)
var señal [2]bool

func proceso0 () {
    for señal[1] == true {
        fmt.Println("Proceso 0: turno: 0")
        fmt.Println("Proceso 0: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    time.Sleep(time.Millisecond * 1000)
    señal[0] = true
    fmt.Println("Proceso 0: Sección Crítica")
    señal[0] = false
    fin <- true
}

func proceso1 () {
    for señal[0] == true {
        fmt.Println("Proceso 1: turno: 0")
        fmt.Println("Proceso 1: {nada}")
        time.Sleep(time.Millisecond * 1000)
    }
    time.Sleep(time.Millisecond * 1000)
    señal[1] = true
    fmt.Println("Proceso 1: Sección Crítica")
    señal[1] = false
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
