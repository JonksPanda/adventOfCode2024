package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func similarityCheck(a int, b []int) int {
	var hits int
	var score int
	hits = 0
	score = 0
	for _, i := range b {
		if a == i {
			hits++
		}
		if i > a {
			break
		}
	}
	score = a * hits
	return score
}

func main() {
	var err error
	var content string
	var lines []string
	var file []byte
	var totalLines int

	file, err = os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var left []int
	var right []int

	// to string
	content = string(file)
	lines = strings.Split(content, "\n")
	totalLines = len(lines)
	for _, line := range lines {
		// 0 = left, 1 = right
		var split []string
		split = strings.Split(line, "   ")

		var rightConv int
		rightConv, err = strconv.Atoi(split[0])
		if err != nil {
			log.Fatal(err)
		}

		var leftConv int
		leftConv, err = strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		left = append(left, rightConv)
		right = append(right, leftConv)
	}
	sort.Ints(left)
	sort.Ints(right)

	var totalDist int

	for i := 0; i < totalLines; i++ {
		var dist int
		dist = diff(left[i], right[i])
		totalDist = totalDist + dist
	}
	println(totalDist)

	var similarityScore int
	for _, i := range left {
		similarityScore = similarityScore + similarityCheck(i, right)
	}
	println(similarityScore)
}
