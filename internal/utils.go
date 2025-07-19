package internal

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func CheckError(err error, readableErrMessage string) {
	if err == nil {
		return
	}

	ThrowError(readableErrMessage + ": " + err.Error())
}

func ThrowError(err string) {
	fmt.Println()
	panic(err)
}

func PrepareColorOutput(output string, color lipgloss.Color) string {
	return lipgloss.NewStyle().Foreground(lipgloss.TerminalColor(color)).Render(output)
}

func CountWords(s string) int {
	s = strings.TrimSpace(s)
	lines := strings.SplitSeq(s, "\n")
	counter := 0

	for line := range lines {
		words := strings.SplitSeq(line, " ")

		for word := range words {
			word = strings.TrimSpace(word)
			if len(word) > 0 {
				counter++
			}
		}
	}

	return counter
}
