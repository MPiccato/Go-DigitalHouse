package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var myArray [5]int
	fmt.Println(unsafe.Sizeof(myArray))
	fmt.Println(len(myArray))
	fmt.Println(cap(myArray))

	myArray = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray)
	myArray[4] = 10
	fmt.Println(myArray)

	myFamily := [4]string{
		"Lorenzo",
		"Lautaro",
		"Cinthia",
		"Martin"}
	fmt.Println(myFamily)
	fmt.Println("Padre:", myFamily[3])

	for i := range myFamily {
		fmt.Println("Índice: ", i, "Familiar: ", myFamily[i])
	}
	for i, value := range myFamily {
		fmt.Println("Índice: ", i, "Familiar: ", value)
	}
}
