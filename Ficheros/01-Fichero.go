package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//01-Fichero.go newtext
	leer_fichero()
	escribir_en_fichero()
	//leer_fichero()
}

func leer_fichero() {
	fmt.Println("Lector")
	texto, err := ioutil.ReadFile("templates/mifichero.txt")
	showError(err)
	fmt.Println(string(texto))
}

func escribir_en_fichero() {
	fmt.Println("Escritura")

	nuevo_texto := os.Args[1] + "\n"
	fmt.Println(string(nuevo_texto))

	fichero, err := os.OpenFile("templates/mifichero.txt", os.O_APPEND, 0777)
	showError(err)

	escribir, err := fichero.WriteString(nuevo_texto)
	showError(err)
	fmt.Println(escribir)

	fichero.Close()

}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
