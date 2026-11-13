package main

import "fmt"

func main() {

	myArray := [5]int{3, 6, 9, 16}
	fmt.Println(myArray, "-longitud:", len(myArray))

	//Los arrays son inmutables, los slices no

	mySlice := []int{} //Así se crea un slice

	//Agreggar datos
	mySlice = append(mySlice, 15, 30, 45)
	fmt.Println(mySlice, "-longitud:", len(mySlice))
	mySlice = append(mySlice, 60, 75, 90)
	fmt.Println(mySlice, "-longitud:", len(mySlice))

	//Agregar datos desde un array
	mySlice2 := myArray[0:3]
	fmt.Println(mySlice2, "-longitud:", len(mySlice2))

	//Si modifico un valor en el array, se modifica en todos los slices que lo referencien
	myArray[0] = 100
	fmt.Println(mySlice2, "-longitud:", len(mySlice2))
}
