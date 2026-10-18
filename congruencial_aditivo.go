package main

import (
	"fmt"
)

func aditivo() []float64 {
	var x, c, m, n int // Declaramos las variables necesarias

	fmt.Print("Ingrese el valor de x0: ")
	fmt.Scan(&x)
	fmt.Print("Ingrese el valor de c: ")
	fmt.Scan(&c)
	fmt.Print("Ingrese el valor de m: ")
	fmt.Scan(&m)

	// Validamos que el valores de x, a, c, m sean mayores que 0
	for x <= 0 {
		fmt.Println("El valor de (x0) debe ser mayor que 0. Por favor, ingresar un numero mayor a 0.")
		fmt.Print("Ingrese el valor de x0: ")
		fmt.Scan(&x)
	}

	for c <= 0 {
		fmt.Println("El valor de (c) debe ser mayor que 0. Por favor, ingresar un numero mayor a 0.")
		fmt.Print("Ingrese el valor de c: ")
		fmt.Scan(&c)
	}

	for m <= 0 {
		fmt.Println("El valor de (m) debe ser mayor que 0. Por favor, ingresar un numero mayor a 0.")
		fmt.Print("Ingrese el valor de m: ")
		fmt.Scan(&m)
	}

	// Validamos que el valor de n sea mayor que 0 y menor o igual a maxN
	for {
		fmt.Printf("Ingrese el valor el limite de numeros generados de 1 a %d: ", maxN)
		fmt.Scan(&n)
		if n > 0 && n <= maxN {
			break
		}
		fmt.Println("Número de valores inválido. Por favor, ingrese un valor entre 1 y", maxN)
	}

	vistos := make(map[int]int) // valor -> posición en la que apareció
	vistos[x] = 0               // Guardamos la posición del primer valor generado
	fmt.Printf("x0 = %d\n", x)

	var sec []float64 // Secuencia normalizada (ri = xi / m)

	// Generamos los números pseudoaleatorios y verificamos si se repiten
	for i := 1; i <= n; i++ {
		x = (x + c) % m
		fmt.Printf("x%d = %d\n", i, x)

		// Verificamos si el número generado ya ha sido visto antes
		if pos, existe := vistos[x]; existe {
			fmt.Printf("\nSe repite: x%d , x%d = %d\n", i, pos, x)
			fmt.Printf("Numeros Generados: %d\n", i)
			break
		}
		vistos[x] = i // Guardamos la posición del número generado
		sec = append(sec, float64(x)/float64(m))
	}
	return sec
}
