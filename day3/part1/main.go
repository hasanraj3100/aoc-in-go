package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data := readFromFile()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(data, -1)

	sum := 0

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	fmt.Println(sum)
}

func readFromFile() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return ""
	}

	return string(data)
}
