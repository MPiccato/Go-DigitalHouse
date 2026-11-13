package function

import (
	"errors"
	"fmt"
)

type Operations int

const (
	op Operations = iota
	SUM
	RE
	DIV
	MULT
)

//Las funciones que empiezan en Mayúscula pueden ser llamadas desde otros paquetes

func Display(myName string) {
	displayNombre(myName)

}

//Las funciones que empiezan en minúsculas solo pueden ser llamadas desde el mismo paquete

func displayNombre(myName string) {
	fmt.Println("Tu nombre es: ", myName)
}

func Suma(x int, y int) int {
	return x + y
}

func Repeat(x int, name string) {
	for i := 0; i < x; i++ {
		fmt.Print(" - ", name)
	}
}

func Calc(operation Operations, a float64, b float64) (float64, error) {
	switch operation {
	case SUM:
		return a + b, nil
	case RE:
		return a - b, nil
	case DIV:
		return a / b, nil
	case MULT:
		return a * b, nil
	default:
		return 0, errors.New("Operacion no valida")
	}
}

func Split(v int) (x, y int) {
	x = 25 * v
	y = v / 2
	return
}

func MSum(values ...float64) float64 {
	var sum float64
	for _, v := range values {
		sum += v
	}
	return sum
}

func MOperations(op Operations, values ...float64) (float64, error) {
	var result float64 = values[0]
	for _, valores := range values[1:] {
		switch op {
		case SUM:
			result += valores
		case RE:
			result -= valores
		case DIV:
			if valores == 0 {
				return 0, errors.New("No se puede dividir por 0")
			}
			result /= valores
		case MULT:
			result *= valores
		default:
			return 0, errors.New("Operacion no valida")
		}

	}
	return result, nil
}
