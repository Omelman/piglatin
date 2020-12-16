package piglatin

import (
	"bufio"
	"fmt"
	"os"
)

func (q *PinLatin) Input() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	q.InputText = text
	return nil
}
