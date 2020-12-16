package main

import (
	"fmt"
	pig "github.com/Omelman/piglatin/piglatin"
	"log"
)

func main() {
	var newPinLatin pig.PinLatin

	err := newPinLatin.Input()
	if err != nil {
		log.Fatalf("Failed to read input: %s", err.Error())
	}

	fmt.Print(newPinLatin.TranslateText())
}
