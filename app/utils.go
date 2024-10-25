package main

import (
	"bufio"
	"log"
	"os"
)

func Log(msg string) {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	resetPrefixLogMessage()

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.Print(msg)
}

func resetPrefixLogMessage() {
	log.SetFlags(log.Flags() &^ log.LstdFlags)
}

func ReadFileEachLine(path string) ([]string, error) {
	f, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanErr := scanner.Err(); scanErr != nil {
		return nil, scanErr
	}

	return lines, nil
}
