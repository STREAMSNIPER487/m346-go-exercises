package main

import (
	"fmt"
	"math"
)

func computeDiscriminant(a, b, c float64) float64 {
	return math.Pow(b, 2) - 4*a*c
}

func computeQuadraticFormula(a, b, c float64) []float64 {
	d := computeDiscriminant(a, b, c)
	if d > 0 {
		root1 := (-b + math.Sqrt(d)) / (2 * a)
		root2 := (-b - math.Sqrt(d)) / (2 * a)
		return []float64{root1, root2}
	} else if d == 0 {
		root := -b / (2 * a)
		return []float64{root}
	}
	return []float64{}
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}
