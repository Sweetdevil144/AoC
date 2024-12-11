package utils

import (
	"log"
	"os"
	"strings"
)

func ParseFile(input string) []string {
	raw, err := os.ReadFile(input)
	if err != nil {
		log.Fatalf("Error Reading File: %v", err)
	}

	data := string(raw)
	arrayData := strings.Split(strings.TrimSpace(data), "\n")
	return arrayData
}