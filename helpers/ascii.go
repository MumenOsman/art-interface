package helpers

import (
	"github.com/common-nighthawk/go-figure"
)

// GenerateASCII converts the input text into an ASCII art banner.
func GenerateASCII(text string) (string, error) {
	myFigure := figure.NewFigure(text, "", false)
	return myFigure.String(), nil
}
