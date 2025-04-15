package main

import (
	"fmt"
	"os"
	"strings"
)

const version = "1.0.0"
const url = "https://github.com/jamieweavis/colstr"
const description = "Simple command line utility to colorize strings based on delimiters contained within them"
const defaultDelimiters = ". / ? = & # :"
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

func colstr(str string, delimiters []string) string {
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

func main() {
	delimiters := strings.Split(defaultDelimiters, " ")
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("colstr v" + version + " (" + colstr(url, delimiters) + ")")
		fmt.Println("\n" + description)
		fmt.Println("\nDefault delimiters: " + defaultDelimiters)
		fmt.Println("\nUsage: colstr <string> [<delimiters>]")
		fmt.Println("\nExamples: colstr 'foo.bar.baz'")
		os.Stdout.WriteString("          colstr 'foo%bar$baz' '%$'\n")
		os.Exit(1)
	}

	str := os.Args[1]
	if len(os.Args) > 2 {
		delimiters = strings.Split(os.Args[2], "")
	}

	result := colstr(str, delimiters)
	fmt.Println(result)
	os.Exit(0)
}
