package main

import "fmt"

func main() {
	saludo("Nico")
	saludo("Roman")
}

func saludo(nombre string) {
	var saludar string = "Hola " + nombre
	if esPar(nombre) {
		saludar = saludar + "!"
	}
	fmt.Println(saludar)
}

func esPar(valor string) (bool){
	return len(valor) % 2 == 0
}
