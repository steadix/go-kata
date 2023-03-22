package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alexsergivan/transliterator"
)

func main() {
	inputFile, err := os.Open("example.txt")
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("example.processed.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	trans := transliterator.NewTransliterator(nil)
	for scanner.Scan() {
		line := scanner.Text()

		// Convert the line to translit
		processedLine := trans.Transliterate(line)

		// Write the transliterated line to the output file
		_, err := fmt.Fprintln(outputFile, processedLine)
		if err != nil {
			panic(err)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
