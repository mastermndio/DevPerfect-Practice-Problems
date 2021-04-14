package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func inputIngestor(filename string) []string {
	file, err := os.Open(filename)

	check(err)

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	file.Close()

	return data
}

func main() {
	ourInput := inputIngestor("input.txt")

	var count int

	for _, letter := range ourInput {
		if letter == "X" {
			count++
		}
	}

	fmt.Println(count)
}
