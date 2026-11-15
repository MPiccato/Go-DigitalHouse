package main

import (
	"fmt"
	"vehicles/vehicles"
)

/*
Las interfaces sirven para manejar datos de distinto tipo
sin modificar el valor de la variable. La "interface{}" es la interfaz
que puede recibir cualquier tipo de dato.
*/

func Display(value interface{}) {
	fmt.Println(value)
}

func main() {

	Display("Martin")
	Display(true)
	Display(48)
	Display(1.60)

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "BIKE"}

	var total float64
	for _, value := range vArray {

		ve, err := vehicles.New(value, 400)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			continue
		}
		total += ve.Distance()
		fmt.Printf("Vehicle: %s Distance: %.2f \n", value, ve.Distance())
	}
	fmt.Println("Total Distance: ", total)

}
