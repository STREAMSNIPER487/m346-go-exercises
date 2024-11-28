package main

import "fmt"

type FullName struct {
	FirstName string
	LastName  string
}

type BirthDate struct {
	Day   int
	Month int
	Year  int
}

type Profile struct {
	Name             FullName
	Birth            BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {

	var me = Profile{
		Name: FullName{
			FirstName: "Arun",
			LastName:  "Aenishänslin",
		},
		Birth: BirthDate{
			Day:   12,
			Month: 6,
			Year:  2007,
		},
		NumberOfSiblings: 1,
		ZodiacSign:       '♊',
	}
	fmt.Println(me)

	fmt.Println("Siblings Before: ", me.NumberOfSiblings)

	me.NumberOfSiblings++

	fmt.Println("Siblings After: ", me.NumberOfSiblings)
}
