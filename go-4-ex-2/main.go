package main

import (
	"fmt"
	"math"
)

func computeHypotenuse(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

type ShortSides struct {
	a float64
	b float64
}

func (s ShortSides) Hypotenuse() float64 {
	return computeHypotenuse(s.a, s.b)
}

func main() {
	fmt.Println("Testing computeHypotenuse:")
	fmt.Printf("Hypotenuse for a=3, b=4: %.2f\n", computeHypotenuse(3, 4))
	fmt.Printf("Hypotenuse for a=5, b=12: %.2f\n", computeHypotenuse(5, 12))
	fmt.Printf("Hypotenuse for a=8, b=15: %.2f\n", computeHypotenuse(8, 15))

	fmt.Println("\nTesting ShortSides.Hypotenuse:")
	sides1 := ShortSides{a: 3, b: 4}
	fmt.Printf("Hypotenuse for a=3, b=4: %.2f\n", sides1.Hypotenuse())

	sides2 := ShortSides{a: 5, b: 12}
	fmt.Printf("Hypotenuse for a=5, b=12: %.2f\n", sides2.Hypotenuse())

	sides3 := ShortSides{a: 8, b: 15}
	fmt.Printf("Hypotenuse for a=8, b=15: %.2f\n", sides3.Hypotenuse())
}
