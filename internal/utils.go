package internal

import (
	"fmt"

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
