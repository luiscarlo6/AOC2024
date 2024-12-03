package utils

import (
	"log"
	"os"
	"strings"
)

func ReadFile(fileName string) []string {
	var lines []string

	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	lines = strings.Split(strings.TrimSpace(string(file)), "\n")

	return lines
}
