package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func medias() []float64 {
	// Pedir la cantidad de números al usuario
	var n int
	for {
		fmt.Print("Ingrese la cantidad de números del 1 al 10: ")
		var s string
		fmt.Scan(&s)
		v, err := strconv.Atoi(s)
		if err == nil && v >= 1 && v <= 10 {
			n = v
			break
		}
		fmt.Println("Entrada inválida. Por favor, ingrese un número del 1 al 10.")
	}

	//Pedir los números al usuario que desea promediar
	fmt.Println("Coloca los números deseados:")
	r := make([]float64, 0, n) // slice para almacenar los números ingresados
	for i := 1; i <= n; {      // bucle hasta que se ingresen n números válidos
		fmt.Printf("%d.- ", i)
		var s string
		fmt.Scan(&s)
		s = strings.Replace(s, ",", ".", 1) // reemplaza la coma por un punto para permitir decimales
		v, err := strconv.ParseFloat(s, 64)
		if err != nil || v < 0 || v > 1 { // verifica que el número esté entre 0 y 1
			fmt.Println("Valor inválido. Ingrese un decimal entre 0 y 1 (ejemplo: 0.58).")
			continue // repite el mismo número sin avanzar
		}
		r = append(r, v) // agrega el número válido al slice
		i++
	}

	// Promedio: suma de los números entre n
	suma := 0.0
	for _, v := range r {
		suma += v
	}
	promedio := suma / float64(n)

	// Z0 = (promedio - 0.5) / (1 / raíz(12n))
	denominador := 1 / math.Sqrt(12*float64(n))
	z0 := (promedio - 0.5) / denominador

	// Valor crítico fijo: z(alfa/2) = 1.96 (alfa = 0.05)
	const zCritico = 1.96

	fmt.Printf("\nx̄ = %.4f\n", promedio)
	fmt.Printf("Z0 = (%.4f - 0.5) / (1/√(12(%d)))\n", promedio, n)
	fmt.Printf("Z0 = (%.4f) / %.4f\n", promedio-0.5, denominador)
	fmt.Printf("Z0 = %.4f\n", z0)

	// Comparación contra ± 1.96
	fmt.Printf("\n-%.2f <= %.4f <= %.2f\n", zCritico, z0, zCritico)
	if z0 >= -zCritico && z0 <= zCritico {
		fmt.Println("Los valores son aceptables: no se rechaza que la media sea 0.5.")
	} else {
		fmt.Println("Los valores no son aceptables: se rechaza que la media sea 0.5.")
	}

	return r
}
