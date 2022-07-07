package main

import "fmt"

func main() {
	//Clase array simple:
	arrays()
	//Clase array multidimensional:
	arrays_multidimensionales()
	//slices:
	slices()
}

func arrays() {
	//Ejemplo 1:
	var peliculas [3]string
	peliculas[0] = "Batman"
	peliculas[1] = "Era del Hielo"
	peliculas[2] = "Harry Potter"

	/*//ejemplo 2:
	peliculas := [3]string{"Batman", "Era del Hielo", "Harry Potter"}*/

	fmt.Println(peliculas)
}

func arrays_multidimensionales() {

	//3 indices, 2 posiciones
	var peliculas [3][2]string
	//indice 1:
	peliculas[0][0] = "Batman"
	peliculas[0][1] = "Era del Hielo"
	//indice 2:
	peliculas[1][0] = "Harry Potter"
	peliculas[1][1] = "Señor de los anillos"
	//indice 3:
	peliculas[2][0] = "Jumanji"
	peliculas[2][1] = "Jurassic Park"

	fmt.Println(peliculas)
}

//Los Slices no tiene el numero de indice declarado:

func slices() {
	peliculas := []string{"Batman", "Era del Hielo", "Harry Potter"}
	//agregar elemento:
	peliculas = append(peliculas, "Pesadilla en elm street")
	fmt.Println(peliculas)
	//cantidad de elementos:
	fmt.Println(len(peliculas))
	//listar desde:
	fmt.Println(peliculas[0:3])
}
