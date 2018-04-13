package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: gendict <dictName> <wordFile> <outFile>\n")
		os.Exit(1)
	}

	dictName := os.Args[1]
	wordFilePath := os.Args[2]
	outFilePath := os.Args[3]

	wordFile, err := os.Open(wordFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer wordFile.Close()

	outFile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	fmt.Fprintf(outFile, "package scrubble\n\n")
	fmt.Fprintf(outFile, "var %s = map[string]bool {\n", dictName)

	scanner := bufio.NewScanner(wordFile)
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Fprintf(outFile, "\t\"%s\": true,\n", word)
	}

	fmt.Fprintf(outFile, "}\n")
}
