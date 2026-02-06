package clipboard

import (
	"github.com/atotto/clipboard"
)

// WriteText writes text to the system clipboard
func WriteText(text string) error {
	return clipboard.WriteAll(text)
}

// ReadText reads text from the system clipboard
func ReadText() (string, error) {
	return clipboard.ReadAll()
}
