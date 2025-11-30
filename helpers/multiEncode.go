package helpers

import "strings"

func MultiLineEncode(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var encodedLines []string

	for _, line := range lines {
		// Use the refactored encoder that doesn't print errors (encoder is less strict)
		encodedLine, _ := SingleLineEncode(line)
		encodedLines = append(encodedLines, encodedLine)
	}

	return strings.Join(encodedLines, "\n"), nil
}
