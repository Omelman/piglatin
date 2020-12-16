package piglatin

import (
	"bufio"
	"fmt"
	"os"
)

func (q *PigLatin) Input() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return text, nil
}
