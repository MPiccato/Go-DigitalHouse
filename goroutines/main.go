// Las goroutines permiten ejecutar procesos en paralelo

package main

import (
	"fmt"
	"time"
)

func main() {
	// se ejecuta la goroutine 1
	go printNames([]string{"Juan", "Maria", "Pedro", "Ana"})
	// se ejecuta la goroutine 2
	go printNumbers()

	time.Sleep(10 * time.Second) // esperamos a que las goroutines terminen
}

func printNames(p []string) {

	for _, names := range p {
		time.Sleep(1 * time.Second)
		fmt.Println(names)
	}
}

func printNumbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
