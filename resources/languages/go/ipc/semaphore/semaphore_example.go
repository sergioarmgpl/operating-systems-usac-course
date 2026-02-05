package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

//var done = make(chan bool)

func main() {
	// Crea un semáforo con capacidad para 2 operaciones concurrentes
	sem := semaphore.NewWeighted(2)
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		go func(id int) {
			// Intenta adquirir un permiso
			if err := sem.Acquire(ctx, 1); err != nil {
				fmt.Printf("Error adquiriendo semáforo: %v\n", err)
				return
			}
                        fmt.Printf("Down Proceso #%d\n",id)
			fmt.Printf("Gorutina %d: Iniciando trabajo...\n", id)
			time.Sleep(2 * time.Second) // Simula trabajo
			fmt.Printf("Gorutina %d: Trabajo terminado.\n", id)
                        fmt.Printf("Up Proceso #%d\n",id)
			defer sem.Release(1) // Asegura liberar el permiso
//                        done <- true
		}(i)
	}
//	<-done
	// Espera un poco para que las gorutinas terminen
	time.Sleep(5 * time.Second)
}
