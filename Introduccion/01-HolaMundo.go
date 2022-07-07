package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Tipos de datos.

	//Variables:
	var suma int = 3 + 1
	var decimal float32 = 3.0 - 1.1
	var booleano bool = true
	var nombre string = "Eduardo Onetto"

	//Constantes:
	const year int = 2022

	//Salidas:
	fmt.Println(suma)
	fmt.Println(decimal)
	fmt.Println(booleano)
	fmt.Println("Hola Mundo desde GO - " + nombre + " - " + strconv.Itoa(year))
}

// --- Tener en consideracion ---
//Correr Codigo GO: go run archivogo.go
//Ver documentacion de librerias:$ godoc time
//Identar y limpiar codigo: go fmt archivogo.go
//Build app to exe: go build archivogo.go
