package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var list []int
	freq := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		value1, err := strconv.Atoi(fields[0])
		value2, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("An error occured while converting inputs to numbers", err)
		}

		list = append(list, value1)
		freq[value2]++
	}

	sum := 0

	for _, num := range list {
		sum += num * freq[num]
	}

	fmt.Println(sum)
}
