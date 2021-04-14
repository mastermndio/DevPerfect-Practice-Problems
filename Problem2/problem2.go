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

	scoreboard := make(map[string]int)

	for _, channel := range ourInput {
		if _, ok := scoreboard[channel]; ok {
			scoreboard[channel]++
			continue
		}

		scoreboard[channel] = 1
	}

	var highest int
	var winner string

	for user, score := range scoreboard {
		if score > highest {
			highest = score
			winner = user
		}
	}

	fmt.Println(winner)
}
