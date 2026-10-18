package main

import (
	"fmt"
	"math"
	"sort"
)

func kolmogorov(r []float64) {
	n := len(r)
	if n == 0 {
		fmt.Println("Primero genere números con las opciones 1 a 3.")
		return
	}

	// Ordenamos una copia para no alterar la secuencia original
	ord := make([]float64, n)
	copy(ord, r)
	sort.Float64s(ord)

	fn := float64(n)
	dMas, dMenos := 0.0, 0.0
	mostrarTabla := n <= 50

	if mostrarTabla {
		fmt.Printf("\n%-4s %-10s %-10s %-10s\n", "i", "r(i)", "D+", "D-")
	}

	for i, v := range ord {
		pos := float64(i + 1)
		dp := pos/fn - v     // i/n - r(i)
		dm := v - (pos-1)/fn // r(i) - (i-1)/n
		if dp > dMas {
			dMas = dp
		}
		if dm > dMenos {
			dMenos = dm
		}
		if mostrarTabla {
			fmt.Printf("%-4d %-10.4f %-10.4f %-10.4f\n", i+1, v, dp, dm)
		}
	}

	d := math.Max(dMas, dMenos)
	fmt.Printf("\nn = %d\n", n)
	fmt.Printf("D+ = %.4f\n", dMas)
	fmt.Printf("D- = %.4f\n", dMenos)
	fmt.Printf("D  = máx(D+, D-) = %.4f\n", d)

	var crit float64
	if n > 35 {
		fmt.Println("Nivel de significancia: 1) 0.10   2) 0.05   3) 0.01")
		fmt.Print("Elija una opción: ")
		var op int
		fmt.Scan(&op)
		switch op {
		case 1:
			crit = 1.22 / math.Sqrt(fn)
		case 3:
			crit = 1.63 / math.Sqrt(fn)
		default:
			crit = 1.36 / math.Sqrt(fn)
		}
		fmt.Printf("Valor crítico = %.4f\n", crit)
	} else {
		fmt.Print("Ingrese el valor crítico D(alfa, n) de la tabla de Kolmogorov-Smirnov: ")
		fmt.Scan(&crit)
	}

	if d <= crit {
		fmt.Println("No se rechaza H0: los números siguen una distribución uniforme.")
	} else {
		fmt.Println("Se rechaza H0: los números no siguen una distribución uniforme.")
	}
}
