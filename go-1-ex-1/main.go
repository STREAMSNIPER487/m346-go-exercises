package main

import "fmt"

var firstName string = "Arun"
var lastName string = "Aenishänslin"

var dayOfBirth int32 = 12
var monthOfBirth int32 = 06
var yearOfBirth int32 = 2007

var numberOfSiblings int32 = 1

var heightInMeters float32 = 1.75

var zodiacSign rune = '\u264A'

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
