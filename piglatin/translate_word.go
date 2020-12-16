package piglatin

import "strings"

const (
	suffix          string = "ay"
	exceptions      string = "aeiou"
	suffixException        = "d" + suffix
)

func (q *PinLatin) TranslateWord(in string) string {
	first := in[0:1]
	if strings.Contains(exceptions, first) {
		return in + suffixException
	} else {
		return in[1:] + first + suffix
	}
}
