package main

import pig "github.com/Omelman/piglatin/piglatin"

type PigLatins interface {
	Input() (string, error)
	TranslateText(in string) pig.PigLatin
	TranslateWord(in string) string
}
