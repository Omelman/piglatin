package piglatin

import (
	"strings"
)

const (
	suffix          string = "ay"
	exceptions      string = "aeiou"
	SuffixException        = "y" + suffix
)

func (q *PigLatin) TranslateWord(in string) string {
	first := in[0:1]
	if strings.Contains(exceptions, strings.ToLower(first)) {
		if strings.ToLower(in[len(in)-1:]) == "y" {
			return in + suffix
		} else {
			return in + SuffixException
		}
	} else {
		var prefix string
		for i, letter := range in {
			if strings.Contains(exceptions, strings.ToLower(string(letter))) {
				prefix = in[:i]
				break
			}
		}
		return in[len(prefix):] + prefix + suffix
	}
}
