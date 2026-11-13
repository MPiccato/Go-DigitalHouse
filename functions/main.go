package main

import (
	"fmt"
	"functions/function"
)

func main() {
	function.Display("Martin")
	function.Display("Lautaro")
	function.Display("Lorenzo")
	function.Display("Cinthia")
	fmt.Printf("La suma es: %d\n", function.Suma(5, 4))
	fmt.Println()
	function.Repeat(5, "Martin")
	fmt.Println()

	result, error := function.Calc(function.SUM, 5.10, 4)
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Resultado:", result)
	}
	multiplicado, errorMult := function.Calc(function.MULT, 4.2, 5)
	if errorMult != nil {
		fmt.Println("Se produjo un error:", errorMult)
	} else {
		fmt.Println("Resultado: ", multiplicado)
	}

	x, y := function.Split(200)
	fmt.Printf("%d , %d", x, y)

	fmt.Printf("\nLa suma es: %.2f", function.MSum(30, 20, 30, 40))

	fmt.Print("\nOperaciones múltiples: \n")

	resultado, error := function.MOperations(function.MULT, 40, 20, 10)
	if error != nil {
		fmt.Println("Error: ", error)
	} else {
		fmt.Println("Resultado: ", resultado)
	}

	resultFactoryFunc := function.FactoryOperation(function.SUM, 10, 20)
	fmt.Println("Resultado: ", resultFactoryFunc(10, 20))
}
