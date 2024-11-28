package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	eyes := rand.Intn(6) + 1
	when := time.Now()

	fmt.Fprintln(os.Stdout, "The dice shows", eyes, "eyes")

	fmt.Fprintln(os.Stderr, "The dice was rolled at", when)

	/*
		go run go-1-ex-3/main.go > eyes.txt 2> dice.log
	*/
}
