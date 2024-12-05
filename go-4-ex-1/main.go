package main

import (
	"errors"
	"fmt"
)

func computeGrade(gotPoints, maxPoints float64) (float64, error) {

	if gotPoints < 0 || maxPoints <= 0 || gotPoints > maxPoints {
		return 0, errors.New("ungültige Punktzahlen: 'gotPoints' darf nicht negativ sein, 'maxPoints' muss positiv sein, und 'gotPoints' darf 'maxPoints' nicht überschreiten")
	}

	grade := 1 + 5*(gotPoints/maxPoints)
	return grade, nil
}

func main() {

	result1, err1 := computeGrade(17.5, 28.0)
	if err1 != nil {
		fmt.Println("Fehler:", err1)
	} else {
		fmt.Printf("Note: %.3f\n", result1)
	}

	result2, err2 := computeGrade(28.0, 28.0)
	if err2 != nil {
		fmt.Println("Fehler:", err2)
	} else {
		fmt.Printf("Note: %.3f\n", result2)
	}

	result3, err3 := computeGrade(5.0, 20.0)
	if err3 != nil {
		fmt.Println("Fehler:", err3)
	} else {
		fmt.Printf("Note: %.3f\n", result3)
	}

	// Test mit ungültigen Werten
	result4, err4 := computeGrade(-5.0, 20.0)
	if err4 != nil {
		fmt.Println("Fehler:", err4)
	} else {
		fmt.Printf("Note: %.3f\n", result4)
	}
}
