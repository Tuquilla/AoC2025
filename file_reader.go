package AoC2025

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(name string) []string {
	var fileInput []string

	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return fileInput
}
