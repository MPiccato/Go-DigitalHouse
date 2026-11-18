package main

import (
	"fmt"
	"os"
)

func main() {

	p := fmt.Println

	file, err := os.Open("main.go")
	if err != nil {
		p("Error opening file: ", err)
		os.Exit(1)
	}
	fmt.Println("File opened successfully")
	v, _ := file.Stat()
	p(v.Name(), v.Size(), v.Mode(), v.ModTime())
	file.Close()
	fmt.Println("File closed successfully")

	data := make([]byte, 1024)

	leerArchivo, err := os.Open("main.go")
	if err != nil {
		p("Error opening file: ", err)
		os.Exit(1)
	}

	c, err := leerArchivo.Read(data)

	p(string(data[:c]), "\nbytes leídos:", c)
	leerArchivo.Close()

	// Genero una variable de entorno
	os.Setenv("MI_ENV_PRUEBA", "MPICCATO DATA")

	// Leo la variable de entorno que acabo de crear
	varEntorno := os.Getenv("MI_ENV_PRUEBA")
	fmt.Println(varEntorno)

	// Dirección en la que me encuentro
	p(os.Getwd())

	//Quiero saber si la variable de entorno existe
	p(os.LookupEnv("MI_ENV_PRUEBA"))

	// Setear datos importantes para la supuesta conexión a una
	// base de datos

	os.Setenv("DB_USERNAME", "MPDATA")
	os.Setenv("DB_PASSWORD", "mysuperpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testing")

	dbURL := os.ExpandEnv("mongodb://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}")
	p(dbURL)
}
