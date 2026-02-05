package main

import (
	"fmt"
//	"os"
	"time"
//	"syscall"
	"github.com/gen2brain/shm"
)

func main() {
	// A unique key for the shared memory segment (e.g., from an ftok call or a fixed number)
	// In a real application, both processes need to agree on this key.
//	key := 12345
//	size := 1024 // Size in bytes
	shmSize := 65536
	
	// Create the shared memory segment with IPC_CREAT flag
	shmId, err := shm.Get(shm.IPC_PRIVATE, shmSize, shm.IPC_CREAT|0777)
//	shmID, err := shm.Get(key, size, syscall.IPC_CREAT|0666)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created shared memory segment with ID: %d\n", shmId)

	// Attach the segment to the process's address space
	data, err := shm.At(shmId, 0, 0)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Attached shared memory at address: %p\n", data)

	// Write data to the shared memory (as a byte slice)
	message := "Hello, shared memory!"
	copy(data, message)
	fmt.Printf("Wrote message: %s\n", message)

	// In a real scenario, you might wait for the other process to finish, 
	// perhaps using semaphores for synchronization.
	fmt.Println("Waiting for 5 seconds for other process to read...")
	time.Sleep(5 * time.Second)

	// Detach the shared memory segment
	if err := shm.Dt(data); err != nil {
		panic(err)
	}
	fmt.Println("Detached shared memory.")

	// Mark the segment for deletion. It will be destroyed after the last detach
	if err := shm.Rm(shmId); err != nil {
		panic(err)
	}
	fmt.Println("Marked shared memory for deletion.")
}
