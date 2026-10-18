package main

import (
	"fmt"
	"math"
)

// Valor crítico aproximado de chi-cuadrada (Wilson-Hilferty)
func chiCritico(v int, z float64) float64 {
	fv := float64(v)
	t := 1 - 2/(9*fv) + z*math.Sqrt(2/(9*fv))
	return fv * t * t * t
}

func chicuadrada(r []float64) {
	n := len(r)
	if n == 0 {
		fmt.Println("Primero genere números con las opciones 1 a 3.")
		return
	}

	var k int
	fmt.Printf("Se generaron %d números. Ingrese el número de intervalos (k): ", n)
	fmt.Scan(&k)
	if k <= 1 {
		fmt.Println("k debe ser mayor que 1.")
		return
	}

	// Frecuencias observadas por intervalo
	obs := make([]int, k)
	for _, v := range r {
		idx := int(v * float64(k))
		if idx >= k {
			idx = k - 1
		}
		obs[idx]++
	}

	// Frecuencia esperada
	esp := float64(n) / float64(k)
	if esp < 5 {
		fmt.Printf("Aviso: la frecuencia esperada es %.2f (< 5). El resultado no es confiable; use menos intervalos o más números.\n", esp)
	}

	// Tabla de resultados
	fmt.Printf("\n%-10s %-18s %-10s %-10s %-12s\n", "Intervalo", "Rango", "Oi", "Ei", "(Ei-Oi)²/Ei")
	chi := 0.0
	for i, o := range obs {
		d := esp - float64(o)
		c := d * d / esp
		chi += c
		ini := float64(i) / float64(k)
		fin := float64(i+1) / float64(k)
		fmt.Printf("%-10d %.4f - %.4f     %-10d %-10.2f %-12.4f\n", i+1, ini, fin, o, esp, c)
	}

	// Grados de libertad
	v := k - 1
	fmt.Printf("\nChi-cuadrada calculada = %.4f\n", chi)
	fmt.Printf("Grados de libertad (k-1) = %d\n", v)

	// Valor crítico según el nivel de significancia
	fmt.Println("Nivel de significancia: 1) 0.10   2) 0.05   3) 0.01")
	fmt.Print("Elija una opción: ")
	var op int
	fmt.Scan(&op)

	// Determinamos el valor crítico aproximado usando la aproximación de Wilson-Hilferty
	var z, alfa float64
	switch op {
	case 1:
		z, alfa = 1.2816, 0.10
	case 3:
		z, alfa = 2.3263, 0.01
	default:
		z, alfa = 1.6449, 0.05
	}

	crit := chiCritico(v, z)
	fmt.Printf("Valor crítico aproximado (alfa = %.2f, %d g.l.) = %.4f\n", alfa, v, crit)
	fmt.Println("(Compárelo con su tabla de chi-cuadrada si necesita el valor exacto.)")

	// Conclusión
	if chi <= crit {
		fmt.Println("No se rechaza H0: los números siguen una distribución uniforme.")
	} else {
		fmt.Println("Se rechaza H0: los números no siguen una distribución uniforme.")
	}
}
