package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {
	var myIntVar int = -12
	fmt.Println("My variable value es ", myIntVar)

	var myUintVar uint = 10
	fmt.Println("My variable value es ", myUintVar)

	var myStringVar string
	myStringVar = "Martin"
	fmt.Println("El valor de mi string es: ", myStringVar)

	var myFloatVar float32 = 12.55
	fmt.Println("El valor de mi float es: ", myFloatVar)

	var myBoolVar bool = true
	fmt.Println("El valor de mi bool es: ", myBoolVar)

	var myOtherVar string = "No puede existir una variable sin instanciar o usar"
	fmt.Println(myOtherVar)

	fmt.Println("Puedo conocer la dirección de memoria de las variables: ", &myOtherVar)

	myVarDirect := 10
	myStringDirect := "Se puede declarar variables de forma directa"
	fmt.Println(myVarDirect, myStringDirect)
	fmt.Println(myStringDirect)

	myBoolDirect := true
	fmt.Println(myBoolDirect)

	var miNombre, miApellido, miEdad string
	miNombre = "Martin"
	miApellido = "Piccato"
	miEdad = "48"
	fmt.Println(miNombre, miApellido, miEdad)

	// Variables Size

	var varS int = 0
	fmt.Println("El tamaño de la variable es: ", varS, "bytes")
	var floatSize32 float32 = 10.5
	fmt.Println("El tamaño de la variable es: ", floatSize32, "bytes")
	var floatSize64 float64 = 10.5
	fmt.Println("El tamaño de la variable es: ", floatSize64, "bytes")

	//El manejo de los flotantes (float32, float64) requiere determinar siempre el tamaño

	var myFloat32Var float32 = 2500.5
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64 = 2500.5
	fmt.Printf("\ntype: %T, value: %f, bytes: %d, bits: %d\n", myFloat64Var, myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	myOtherFloat := 5500.34 // Por defecto eligirá la capacidad de la computadora, 64 bits en este caso
	fmt.Printf("type: %T, value: %f, bytes: %d, bits: %d\n", myOtherFloat, myOtherFloat, unsafe.Sizeof(myOtherFloat), unsafe.Sizeof(myOtherFloat)*8)

	// Cadena de texto
	var myStringText string = "Hola Go"
	fmt.Println("\nMi variable es: ", myStringText)
	fmt.Println("\nMuestra el código ASCII de: ", myStringText[0])

	var myStringMultiple string = `Esta es una variable con 
	muchos saltos de lineas
	que se hacen automáticamente`
	fmt.Println(myStringMultiple)

	// casteo de variables
	{
		floatVariable := 55.2
		fmt.Printf("\ntype: %T, value: %.3f", floatVariable, floatVariable)

		floatStringVar := fmt.Sprintf("%.2f", floatVariable)
		fmt.Printf("\ntype: %T, value: %s\n", floatStringVar, floatStringVar)

		intVar := 20
		fmt.Printf("type: %T, value: %d\n", intVar, intVar)
		intStringVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStringVar, intStringVar)

		intVar2 := 30
		strConvertida := strconv.Itoa(intVar2)
		fmt.Printf("type: %T, value: %d\n", intVar2, intVar2)
		fmt.Printf("type: %T, value: %s", strConvertida, strConvertida)

	}

	{
		strToInt := "30"
		fmt.Printf("type: %T, value: %s \n", strToInt, strToInt)

		intStrToInt, err := strconv.Atoi(strToInt)
		fmt.Printf("type: %T, value: %d, error: %v\n", intStrToInt, intStrToInt, err)

		intToStr2, _ := strconv.ParseInt("25", 10, 64)
		fmt.Printf("type: %T, value: %d\n", intToStr2, intToStr2)

	}
}
