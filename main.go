package main

import (
	"fmt"
)

const maxN = 4000 // Definimos una constante para el límite máximo de números generados

var ultima []float64 // última secuencia generada

// Submenú: el usuario escoge qué prueba de uniformidad aplicar
func uniformidad(r []float64) {
	if len(r) == 0 {
		fmt.Println("Primero genere números con las opciones 1 a 3.")
		return
	}

	fmt.Println("\nPrueba de uniformidad")
	fmt.Println("---------------------")
	fmt.Println("1. Chi-cuadrada")
	fmt.Println("2. Kolmogorov-Smirnov")
	fmt.Print("Elija la prueba: ")

	var op int
	fmt.Scan(&op)

	switch op {
	case 1:
		chicuadrada(r)
	case 2:
		kolmogorov(r)
	default:
		fmt.Println("Opción inválida.")
	}
}

func main() {
	for {
		fmt.Println("\nGeneradores Congruenciales")
		fmt.Println("--------------------------")
		fmt.Println("Elija una opción:")
		fmt.Println("1. Congruencial lineal")
		fmt.Println("2. Congruencial multiplicativo")
		fmt.Println("3. Congruencial aditivo")
		fmt.Println("4. Prueba de uniformidad")
		fmt.Println("5. Pruebas de Medias")
		fmt.Println("6. Salir")

		var opcion int
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			ultima = lineal()
		case 2:
			ultima = multiplicativo()
		case 3:
			ultima = aditivo()
		case 4:
			uniformidad(ultima)
		case 5:
			medias()
		case 6:
			fmt.Println("Hasta luego.")
			return
		default:
			fmt.Println("Opción inválida. Elija un número del 1 al 6.")
		}
	}
}
