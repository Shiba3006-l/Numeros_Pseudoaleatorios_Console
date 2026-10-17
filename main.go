package main

import (
	"fmt"
)

const maxN = 4000 // Definimos una constante para el límite máximo de números generados

func main() {
	for {
		fmt.Println("\n===== Generadores Congruenciales =====")
		fmt.Println("1. Congruencial lineal")
		fmt.Println("2. Congruencial multiplicativo")
		fmt.Println("3. Congruencial aditivo")
		fmt.Println("4. Salir")
		fmt.Print("Elija una opción: ")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			lineal()
		case 2:
			multiplicativo()
		case 3:
			aditivo()
		case 4:
			fmt.Println("Hasta luego.")
			return
		default:
			fmt.Println("Opción inválida. Elija un número del 1 al 4.")
		}
	}
}
