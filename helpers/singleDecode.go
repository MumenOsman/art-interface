package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

func SingleDecode(input string) (string, error) {

	var result strings.Builder

	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			start := i

			j := strings.Index(input[i+1:], "]")

			if j == -1 {
				return "", fmt.Errorf("Error")
			}

			j += i + 1 // j is now the absolute index of ']' not just from the '[' 

			contentString := input[start+1 : j] // does not include the j 
			spaceIndex := strings.Index(contentString, " ")

			if spaceIndex == -1 || spaceIndex == 0 || spaceIndex == len(contentString)-1 { // it didnt exist or it was first or it was last
				return "", fmt.Errorf("Error")
			}

			countStr := contentString[:spaceIndex]
			charsToRepeat := contentString[spaceIndex+1:]

			counter, err := strconv.Atoi(countStr)

			if err != nil || counter < 0 {
				return "", fmt.Errorf("Error")
			}

			result.WriteString(strings.Repeat(charsToRepeat, counter))

			i = j + 1
		} else {
			result.WriteByte(input[i])
		}
	}
	return result.String(), nil
}
