package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	// fecha := time.Date(2024, time.October, 14, 19, 20, 50, 0, time.Local)
	// p(fecha)

	// Si quiero agregar 7 días
	fecha := time.Now()
	p(fecha)

	fecha = fecha.Add(time.Hour * 24 * 7)
	p(fecha)

	p(fecha.Day())
	p(fecha.Month())
	p(fecha.Year())
	p(fecha.Hour())
	p(fecha.Minute())
	p(fecha.Second())
	p(fecha.Nanosecond())

	hoy := time.Now()
	p("La fecha actual es vieja: ", hoy.Before(hoy))
	p("La fecha actual es nueva: ", hoy.After(hoy))
	p("La fecha actual es igual: ", hoy.Equal(hoy))

	fechaAnterior := time.Date(2024, time.October, 14, 19, 20, 50, 0, time.Local)

	diff := fecha.Sub(fechaAnterior)

	p(diff)

}
