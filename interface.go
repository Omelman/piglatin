package main

type PigLatins interface {
	Input() error
	TranslateText() string
	TranslateWord(in string) string
}
