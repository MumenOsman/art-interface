package helpers

import (
	"fmt"
	"strings"
)

func SingleLineEncode(artString string) (string, error) {
	var encoded strings.Builder
	n := len(artString)
	i := 0

	for i < n {
		maxCount := 1
		bestPattern := ""
		bestLength := 1

		// Search for the longest repeating pattern starting at i
		// Limit the pattern length L to a reasonable max (e.g., n/2 or 10)
		searchLimit := (n - i) / 2
		if searchLimit > 10 {
			searchLimit = 10
		}

		for L := 1; L <= searchLimit; L++ {
			if i+L > n {
				break
			}
			pattern := artString[i : i+L]
			currentCount := 1
			nextIndex := i + L

			for nextIndex+L <= n && artString[nextIndex:nextIndex+L] == pattern {
				currentCount++
				nextIndex += L
			}

			// If the compression is effective (saves characters), record it
			// Condition: len("[N P]") < N * len(P)
			// A minimal effective compression is 2x1 -> [2 1], saves 1 char.
			// A minimal effective compression is 3x1 -> [3 1], saves 2 chars.
			// The current check implicitly favors the longest match.
			if currentCount > maxCount {
				maxCount = currentCount
				bestPattern = pattern
				bestLength = L
			}
		}

		if maxCount > 1 {
			// Found a repeatable pattern: compress it
			// Ensure the pattern doesn't contain the reserved characters '[' and ']'
			// Ensure the pattern doesn't contain the reserved character ']'
			// We can allow '[' in the pattern.
			if strings.Contains(bestPattern, "]") {
				// Fallback to single character if pattern contains reserved chars
				if artString[i] == '[' {
					encoded.WriteString("[1 []")
				} else {
					encoded.WriteByte(artString[i])
				}
				i++
			} else {
				encoded.WriteString(fmt.Sprintf("[%d %s]", maxCount, bestPattern))
				i += (maxCount * bestLength)
			}
		} else {
			// No significant compression found: append the single character
			// We check for reserved characters, but they should be printed as-is
			if artString[i] == '[' {
				encoded.WriteString("[1 []")
			} else {
				encoded.WriteByte(artString[i])
			}
			i++
		}
	}

	return encoded.String(), nil
}
