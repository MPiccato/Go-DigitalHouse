package main

import "fmt"

func main() {

	yearsOld := 48
	fmt.Println(yearsOld < 48)  // false
	fmt.Println(yearsOld > 48)  // false
	fmt.Println(yearsOld <= 48) // true
	fmt.Println(yearsOld >= 48) // true
	fmt.Println(yearsOld == 48) // true
	fmt.Println(yearsOld != 48) // false

	fmt.Println("\nUso del operador OR\n")

	fmt.Println((yearsOld < 48) || (yearsOld > 48))  // false
	fmt.Println((yearsOld < 48) || (yearsOld == 48)) // true
	fmt.Println((yearsOld < 48) || (yearsOld > 48))  // false
	fmt.Println((yearsOld < 48) || (yearsOld > 48))  // false

	fmt.Println("\nUso del operador AND\n")

	fmt.Println((yearsOld < 50) && (yearsOld == 48)) // true

	fmt.Println((yearsOld < 50) && (yearsOld == 40)) // false

	fmt.Println((yearsOld > 50) && (yearsOld == 48)) // false

	fmt.Println("\nOperadores de Asignación\n")

}
