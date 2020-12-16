package main

import (
	"fmt"
	pig "github.com/Omelman/piglatin/piglatin"
	"log"
)

func main() {
	var newPinLatin pig.PigLatin

	input, err := newPinLatin.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}

	newPinLatin = newPinLatin.TranslateText(input)
	fmt.Print(newPinLatin)
}
