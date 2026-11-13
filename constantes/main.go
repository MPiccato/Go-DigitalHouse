package main

import "fmt"

func main() {
	const pi = 3.1415
	fmt.Println("El valor de pi es: ", pi)

	const myStringDirect string = "No conviene declarar constantes de forma directa"
	fmt.Println(myStringDirect)

	myVarDirect := 10
	fmt.Println(myVarDirect)

	myScopeConst := 20
	{
		myScopeConst := 10
		fmt.Println("Dentro del Scope: ", myScopeConst)
	}
	fmt.Println("Fuera del Scope: ", myScopeConst)

}
