package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func convertCelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func convertFahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (c Celsius) ConvertToFahrenheit() Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func (f Fahrenheit) ConvertToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	var c1, c2, c3 Celsius = 0, 25, 100
	f1 := convertCelsiusToFahrenheit(c1)
	f2 := convertCelsiusToFahrenheit(c2)
	f3 := convertCelsiusToFahrenheit(c3)
	fmt.Printf("%v °C = %v °F\n", c1, f1)
	fmt.Printf("%v °C = %v °F\n", c2, f2)
	fmt.Printf("%v °C = %v °F\n", c3, f3)

	reverseC1 := convertFahrenheitToCelsius(f1)
	reverseC2 := convertFahrenheitToCelsius(f2)
	reverseC3 := convertFahrenheitToCelsius(f3)
	fmt.Printf("%v °F = %v °C\n", f1, reverseC1)
	fmt.Printf("%v °F = %v °C\n", f2, reverseC2)
	fmt.Printf("%v °F = %v °C\n", f3, reverseC3)

	cozy := Celsius(23.0)
	cold := Fahrenheit(15.3)
	fmt.Printf("%v °C = %v °F\n", cozy, cozy.ConvertToFahrenheit())
	fmt.Printf("%v °F = %v °C\n", cold, cold.ConvertToCelsius())

	fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(23.0)))
	var c Celsius = 23
	fmt.Println(c.ConvertToFahrenheit().ConvertToCelsius())
}
