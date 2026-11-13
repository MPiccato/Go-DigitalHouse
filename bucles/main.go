package main

import "fmt"

func main() {

	// Bucle for
	{
		sum := 0
		// Se declara la variable que vamos a usar en el bucle
		for i := 0; i <= 10; i++ {

			sum += i
		}
		fmt.Println(sum)

		sum = 1

		for {
			if sum > 100 {
				break
			}
			sum++
		}
		fmt.Println(sum)
	}
}
