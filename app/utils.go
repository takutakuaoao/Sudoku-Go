package app

import (
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
