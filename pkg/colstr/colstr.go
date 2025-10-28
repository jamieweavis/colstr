// Package colstr provides functionality to colorize strings based on delimiters.
package colstr

import "strings"

// Color constants for terminal output
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

// DefaultDelimiters is the default set of delimiters used for string colorization
const DefaultDelimiters = ". / ? = & # :"

// Colorize takes a string and a slice of delimiters, and returns a colorized version
// of the string where each segment between delimiters is colored differently.
// The delimiters themselves are colored in white.
func Colorize(str string, delimiters []string) string {
	colors := []string{Red, Cyan, Yellow, Green, Blue, Purple}
	colorIndex := 0
	pos := 0
	result := ""

	for pos < len(str) {
		nextSepPos := len(str)
		nextSep := ""

		for _, sep := range delimiters {
			if idx := strings.Index(str[pos:], sep); idx >= 0 && pos+idx < nextSepPos {
				nextSepPos = pos + idx
				nextSep = sep
			}
		}

		if nextSepPos > pos {
			part := str[pos:nextSepPos]
			result += colors[colorIndex] + part + Reset
			colorIndex = (colorIndex + 1) % len(colors)
		}

		if nextSepPos < len(str) {
			result += White + nextSep + Reset
			pos = nextSepPos + len(nextSep)
		} else {
			break
		}
	}

	return result
}
