package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var path string
	fmt.Scan(&path)

	lines, err := ReadFileEachLine(path)

	if err != nil {
		fmt.Printf("Error!\n%v\n", err)
		os.Exit(1)
	}

	allNumbers := parse(lines)

	showQuestion(allNumbers)

	s := time.Now()
	result := NewResolver(*NewBoard(allNumbers)).Resolve()
	end := time.Since(s)

	showAnswer(result.board.rows)
	fmt.Printf("\nprocess time: %s\n\n", end)
}

func parse(lines []string) [9][9]uint8 {
	allNumbers := [9][9]uint8{}

	for rowKey, line := range lines {
		rowNumbers := [9]uint8{}

		for columnKey, number := range strings.Fields(line) {
			n, _ := strconv.Atoi(number)
			rowNumbers[columnKey] = uint8(n)
		}

		allNumbers[rowKey] = rowNumbers
	}

	return allNumbers
}

func showQuestion(question [9][9]uint8) {
	fmt.Printf("\nQuestion...\n=====================\n")

	for _, row := range question {
		fmt.Println(row)
	}
}

func showAnswer(result [9][9]uint8) {
	fmt.Printf("\nAnswer:)\n=====================\n")
	for _, row := range result {
		fmt.Println(row)
	}
}
