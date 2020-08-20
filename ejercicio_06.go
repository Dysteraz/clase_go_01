package main

func main() {
	print(reemplazarPalabra("Hola soy homero simpson", "simpson", "thomsom"))
}

func reemplazarPalabra(frase string, palabra1 string, palabra2 string) string {
	largoPalabra := len(palabra1)
	for i := 0; i <= len(frase)-len(palabra1)-1; i++ {
		if frase[i:largoPalabra] == palabra1[:] {
			for j := 1; j <= largoPalabra; j++ {
				append(frase[i], "")
			}
			append(frase[i], palabra2[:])
			return frase
		}
		largoPalabra++
	}
	return "Frase no encontrada"
}
