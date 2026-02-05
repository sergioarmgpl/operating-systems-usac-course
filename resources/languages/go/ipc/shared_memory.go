package main

import (
    "fmt"
    "syscall"
)

const SHMSIZE = 1024

func main() {
    // Create a shared memory segment
    shmid, err := syscall.shmget(syscall.KEY_IPC_PRIVATE, SHMSIZE, 0666|syscall.IPC_CREAT)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("shmid%d",shmid)
    // Attach to the shared memory segment
    shmptr, err := syscall.Shmat(shmid, 0, 0)
    if err != nil {
        fmt.Println(err)
        return
    }

    // Write to the shared memory segment
    message := []byte("Hello from process 1!")
    copy(shmptr, message)

    // Detach from the shared memory segment
    err = syscall.Shmdt(shmptr)
    if err != nil {
        fmt.Println(err)
        return
    }
}
