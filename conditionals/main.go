package main

import "fmt"

func main() {

	{
		yearsOld := 48

		if yearsOld < 18 {
			fmt.Printf("Eres menor de edad, tienes %d años", yearsOld)
		} else {
			fmt.Printf("Eres mayor de edad, tienes %d años", yearsOld)
		}
	}
	{
		boolVar := true
		if boolVar {
			fmt.Println("\nLa variable es verdadera")
		} else {
			fmt.Println("\nLa variable es falsa")
		}
	}

	// Sentencia IF
	{
		if value := true; !value {
			fmt.Println("\nLa variable es falsa")
		} else {
			fmt.Println("\nLa variable es verdadera")
		}
	}
	// Sentencia Switch
	{
		number := 3

		switch number {
		case 1:
			fmt.Println("El numero es 1")
		case 2:
			fmt.Println("El numero es 2")
		case 3:
			fmt.Println("El numero es 3")
		default:
			fmt.Println("El numero no es 1,2,3")
		}

		//En caso de usar funciones, podemos generar una variable dentro del mismo switch
		switch value := 1; value {
		case 1:
			fmt.Println("El numero es 1")
		case 2:
			fmt.Println("El numero es 2")
		case 3:
			fmt.Println("El numero es 3")
		default:
			fmt.Println("El numero no es 1,2,3")
		}

		switch {
		case number < 1:
			fmt.Println("El numero es menor a 1")
		case number > 1:
			fmt.Println("El numero es mayor a 1")
		}

	}

}
