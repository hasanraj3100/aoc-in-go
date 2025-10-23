package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the file")
		return
	}

	defer file.Close()

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			xStr, yStr := match[1], match[2]

			x, errX := strconv.Atoi(xStr)
			if errX != nil {
				fmt.Println("Error converting x string to int")
				continue
			}
			y, errY := strconv.Atoi(yStr)
			if errY != nil {
				fmt.Println("Error converting y string to int")
				continue
			}
			sum += x * y

		}

	}

	fmt.Println(sum)
}
