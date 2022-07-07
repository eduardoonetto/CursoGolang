package main

import "fmt"

func main() {
	num1 := 5
	num2 := 10
	result := suma(num1, num2)
	//ejemplo 1:
	fmt.Println(dosretornos_2())
	//ejemplo 2:
	fmt.Println("El Resultado es:")
	fmt.Println(result)
	//ejemplo 3:
	fmt.Println(gorras(21, "CLP"))
	//ejemplo 4:
	multiparams("Hola", "CLP", "Rojo", "Algo")

}

func suma(numero1 int, numero2 int) int {
	var result int = numero1 + numero2
	return result
}

func dosretornos() (string, string) {
	return "Hola", "Alumno"
}

func dosretornos_2() (dato1 string, dato2 string) {
	dato1 = "Hola"
	dato2 = "Alumno"
	//ya le dimos los params en la funcion, no es necesario especificarlo.
	return
}

//Cluster
func gorras(pedido float32, moneda string) (string, float32, string) {

	const precio_unit float32 = 3500
	//funcion anonima:
	total := func() float32 {
		return pedido * 3500
	}
	return "El precio del pedido: $", total(), moneda
}

func multiparams(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}
