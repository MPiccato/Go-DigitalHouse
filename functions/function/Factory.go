package function

func FactoryOperation(op Operations, x, y float64) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sumar
	case RE:
		return restar
	case MULT:
		return multiplicar
	case DIV:
		return dividir
	default:
		return nil
	}

}

func sumar(a, b float64) float64 {
	return a + b
}

func restar(a, b float64) float64 {
	return a - b
}

func multiplicar(a, b float64) float64 {
	return a * b
}

func dividir(a, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
