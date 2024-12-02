package utils

import (
	"log"
	"os"
	"strings"
)

func GetFileRows(filePath string) []string {
	var err error
	var file []byte
	var content string
	var lines []string

	file, err = os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	content = string(file)

	lines = strings.Split(content, "\n")

	return lines
}
