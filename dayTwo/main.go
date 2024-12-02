package main

import (
	"github.com/JonksPanda/adventOfCode2024/utils"
	"strings"
	"strconv"
	"log"
)

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func IsSorted(slice []int) bool {
	var asc bool
	var desc bool

	asc = true
	desc = true

	for i := 0; i < len(slice); i++ {
		if i+1 < len(slice) {
			if slice[i] < slice[i+1] {
				desc = false
			} else if slice[i] > slice[i+1] {
				asc = false
			}
		}
	}
	if !asc && !desc {
		return false
	} else {
		return true
	}
}

func main() {
	var lines []string
	var err error
	var safeReports int

	lines = utils.GetFileRows("input.txt")

	
	for _, line := range lines {
		var isSafe bool
		var levels []int
		isSafe = true

		var split []string
		split = strings.Split(line, " ")

		// to int list
		for _, levelString := range split {
			var level int
			level, err = strconv.Atoi(levelString)
			if err != nil {
				log.Fatal(err)
			}

			levels = append(levels, level)
		}

		if !IsSorted(levels) {
			isSafe = false
		}

		for i := 0; i < len(levels); i++ {
			if i+1 < len(levels) {
				var diff int
				diff = Diff(levels[i], levels[i+1])
				if diff > 3 || diff < 1 {
					isSafe = false
				}
			}
		}

		if isSafe {
			safeReports = safeReports + 1
		}
	}
	println(safeReports)
}