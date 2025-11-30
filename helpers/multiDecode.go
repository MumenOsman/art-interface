package helpers

import "strings"

func MultiDecode(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var decodedLines []string

	for _, line := range lines {
		decodedLine, err := SingleDecode(line)
		if err != nil {
			return "", err // Abort and return "Error" on the first malformed line
		}
		decodedLines = append(decodedLines, decodedLine)
	}

	return strings.Join(decodedLines, "\n"), nil
}
