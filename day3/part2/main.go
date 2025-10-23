package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file", err)
		return
	}

	fmt.Println(calculateString(string(data)))
}

func calculateString(line string) int {
	sum := 0
	enabled := true
	re := regexp.MustCompile(`don't\(\)|do\(\)|mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(line, -1)

	for _, match := range matches {
		switch match[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				x, _ := strconv.Atoi(match[1])
				y, _ := strconv.Atoi(match[2])
				sum += x * y
			}
		}
	}

	return sum
}
