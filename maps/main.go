package main

import "fmt"

func main() {

	familia := make(map[string]string)
	familia["padre"] = "Martin"
	familia["madre"] = "Cinthia"
	familia["hijoMayor"] = "Lautaro"
	familia["hijoMenor"] = "Lorenzo"

	fmt.Println(familia)
	fmt.Println(familia["hijoMenor"])

	map2 := make(map[int][]string)
	map2[1] = []string{"Juan", "Pedro", "Lucia", "Pedro"}
	map2[2] = []string{"2", "34", "56", "4000"}

	for i, value := range map2 {
		fmt.Println(i, value)
	}
	for _, value := range map2 {
		for _, value := range value {
			fmt.Println(value)
		}
	}
	fmt.Println(map2[1][2])

	fmt.Println("Puedo saber si el valor existe o no...")
	map3 := make(map[int]string)
	map3[1] = "Lucas"
	map3[2] = "Juan"
	map3[3] = "Rocio"

	valor, ok := map3[4]
	fmt.Println("El valor map3,4: ", valor, " ¿Existe?: ", ok)

	map3[4] = " "
	valor, ok = map3[4]
	fmt.Println("El valor map3,4: ", valor, " ¿Existe?: ", ok)

}
