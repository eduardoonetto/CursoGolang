package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	var miGorra = Gorra{
		marca:  "Nike",
		color:  "Negro",
		precio: 3500.00,
		plana:  false}

	fmt.Println(miGorra)
}
