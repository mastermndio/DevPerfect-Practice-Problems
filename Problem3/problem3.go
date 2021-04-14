package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	rubric := map[string]int{"!hate": -2, "!dislike": -1, "!like": 1, "!love": 2}

	scoreboard := make(map[string]int)
	var splitInput []string

	for _, channel := range ourInput {
		splitInput = strings.Split(channel, " ")

		if _, ok := scoreboard[splitInput[1]]; ok {
			scoreboard[splitInput[1]] += rubric[splitInput[0]]
			continue
		}

		scoreboard[splitInput[1]] += rubric[splitInput[0]]
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
