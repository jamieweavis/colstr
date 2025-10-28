package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jamieweavis/colstr/pkg/colstr"
)

const (
	version     = "1.0.0"
	url         = "https://github.com/jamieweavis/colstr"
	description = "Simple command line utility to colorize strings based on delimiters contained within them"
)

func main() {
	delimiters := strings.Split(colstr.DefaultDelimiters, " ")
	if len(os.Args) < 2 || len(os.Args) > 3 {
		printUsage(delimiters)
		os.Exit(1)
	}

	str := os.Args[1]
	if len(os.Args) > 2 {
		delimiters = strings.Split(os.Args[2], "")
	}

	result := colstr.Colorize(str, delimiters)
	fmt.Println(result)
	os.Exit(0)
}

func printUsage(delimiters []string) {
	fmt.Println("colstr v" + version + " (" + colstr.Colorize(url, delimiters) + ")")
	fmt.Println("\n" + description)
	fmt.Println("\nDefault delimiters: " + colstr.DefaultDelimiters)
	fmt.Println("\nUsage: colstr <string> [<delimiters>]")
	fmt.Println("\nExamples: colstr 'foo.bar.baz'")
	os.Stdout.WriteString("          colstr 'foo%bar$baz' '%$'\n")
}
