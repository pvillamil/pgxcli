package ui

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/muesli/termenv"
)

//go:embed ascii.txt
var asciiArt string

func orcaStr(out *termenv.Output) string {
	oceanBlue := out.Color("#4E7080")
	steel := out.Color("#A8BDC8")

	// Strip leading blank lines and trailing spaces per line.
	lines := strings.Split(asciiArt, "\n")
	for len(lines) > 0 && strings.TrimSpace(lines[0]) == "" {
		lines = lines[1:]
	}
	for len(lines) > 0 && strings.TrimSpace(lines[len(lines)-1]) == "" {
		lines = lines[:len(lines)-1]
	}
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " ")
	}

	var sb strings.Builder
	for _, r := range strings.Join(lines, "\n") {
		switch r {
		case ':', ';', '.', ',', '-':
			sb.WriteString(out.String(string(r)).Foreground(oceanBlue).String())
		case '▆', '▀':
			sb.WriteString(out.String(string(r)).Foreground(steel).String())
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// orcaView returns the colored orca art for use in TUI layouts.
func orcaView() string {
	return orcaStr(termenv.NewOutput(os.Stdout))
}

// PrintBanner prints the colored ASCII art banner and a welcome line.
func PrintBanner(version string) {
	out := termenv.NewOutput(os.Stdout)
	green := out.Color("#02BF87")

	fmt.Print(orcaStr(out))
	fmt.Printf("\n  %s  %s\n\n",
		out.String("pgxcli v"+version).Foreground(green).Bold().String(),
		out.String("\\q to quit").Foreground(out.Color("240")).String(),
	)
}
