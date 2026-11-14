package main

import "fmt"

type MyStruct struct {
	ID   int
	Name string
}

func (m MyStruct) Set(name string) {
	fmt.Printf("Address of parameter 'm': %p\n", &m)
	m.Name = name
}
func (m *MyStruct) SetP(name string) {
	fmt.Printf("Address of parameter 'm': %p\n", m)
	m.Name = name
}

func main() {
	var i int
	var iP *int

	i = 50
	iP = &i

	fmt.Println("El valor de i es", i)
	fmt.Println("La dirección de memoria de i es", &i)
	fmt.Println("El valor de iP es", *iP)
	fmt.Println("La dirección de memoria de iP es", &iP)

	*iP = 100
	fmt.Println("El valor de i es", i)
	fmt.Println("La dirección de memoria de i es", &i)
	fmt.Println("El valor de iP es", *iP)
	fmt.Println("La dirección de memoria de iP es", &iP)

	miEstructura := &MyStruct{ID: 1, Name: "Hola"}
	miEstructura.SetP("Adios")
	fmt.Printf("La direccion es: %p\n", miEstructura)
	fmt.Printf("Elementos: %d - %s", miEstructura.ID, miEstructura.Name)

}
