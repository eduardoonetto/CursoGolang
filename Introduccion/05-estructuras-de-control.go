package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("*** Mi Programa con GO ***")

	// Condicional if:
	//condicional_if()

	//Bucle for:
	//bucle_for()

	//Switch:
	usar_switch()
}

func condicional_if() {
	//get params from console:
	params := os.Args
	edad, _ := strconv.Atoi(params[1])

	//Usar && - || (and - or)
	//Anidado usar }else if condition {}
	if edad >= 18 && edad <= 99 {
		fmt.Println("Eres mayor de edad.")
	} else {
		fmt.Println("Eres menor de edad o demasiado mayor.")
	}
}

func bucle_for() {

	for i := 0; i < 3; i++ {
		fmt.Println(i, "iteracion")
	}

	peliculas := []string{"Batman", "Era del Hielo", "Harry Potter"}

	for i := 0; i < len(peliculas); i++ {
		fmt.Println(i, peliculas[i])
	}

	//foreach
	for _, pelicula := range peliculas {
		fmt.Println(pelicula)
	}

}

func usar_switch() {
	hoy := time.Now().Weekday()

	switch hoy {
	case 0:
		fmt.Println("Domingo")
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	default:
		fmt.Println("Hoy es otro dia de la semana")

	}
}
