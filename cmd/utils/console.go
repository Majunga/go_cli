package utils

import (
	"bufio"
	"os"
	"strings"
)

func Read() (string, error) {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString('\n')

		if len(strings.TrimSpace(text)) > 0 {
			return strings.TrimSpace(text), err
		}
	}
}
