package main

import "fmt"

func main() {
	saludo("Nico")
}

func saludo(nombre string) {
	var saludo string = "Hola " + nombre
	if esPar(nombre) {
		saludo = saludo + "!"
	}
	fmt.Println(saludo)
}

func esPar(valor string) (bool){
	if len(valor) % 2 == 0 {
		return true
	}
	return false
}
