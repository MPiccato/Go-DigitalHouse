package main

import (
	"encoding/json"
	"fmt"
	"structs-app/commerce"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

// Métodos de la estructura User

// Display es un método que muestra la estructura en formato plano, bytes y string json
func (us User) Display() {
	userJson, _ := json.Marshal(us)
	fmt.Println("\nEn formato bytes:\n")
	fmt.Println(userJson)
	fmt.Print()
	fmt.Println("\nEn formato json string:\n")
	fmt.Println(string(userJson))

}

//Para modificar la estructura original debo agregar el * en el nombre de la estructura

func (u *User) SetName(name string) {
	u.Name = name

}

func main() {

	user := User{
		ID:       "1",
		Name:     "Martin Piccato",
		Email:    "mpiccato@gmail.com",
		IsActive: true,
	}

	fmt.Print()
	user.Display()
	user.SetName("Juan Piccato")
	user.Display()

	// Uso del paquete commerce

	p1 := commerce.Product{
		Name: "Heladera",
		Type: commerce.Type{
			Code:        "A",
			Description: "Linea Blanca",
		},
		Price:    200000,
		Quantity: 5,
		Tags:     []string{"heladera", "electrodoméstico", "freezer"},
	}
	p2 := commerce.Product{
		Name: "TV",
		Type: commerce.Type{
			Code:        "E",
			Description: "Entretenimiento",
		},
		Price:    800000,
		Quantity: 1,
		Tags:     []string{"tv", "electrodoméstico", "televisor"},
	}

	carrito := commerce.NewCart(1)
	carrito.AddProduct(p1, p2)
	fmt.Println(int(carrito.Total()))

}
