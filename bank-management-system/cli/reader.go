package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadInput(prompt string) (string, error) {
	fmt.Print("\n", prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil
}
