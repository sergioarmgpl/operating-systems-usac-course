package main
/* producer-consumer problem in Go */


import ("fmt")

var done = make(chan bool)
var msgs = make(chan int)

func produce () {
    for i := 0; i < 10; i++ {
        msgs <- i
    }
    done <- true
}

func consume () {
    for {
      msg := <-msgs
      fmt.Println(msg)
   }
}

func main () {
   go produce()
   go consume()
   <- done
}
