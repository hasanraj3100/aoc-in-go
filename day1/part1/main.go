package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var list1, list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		value1, err := strconv.Atoi(fields[0])
		value2, err := strconv.Atoi(fields[1])
		if err != nil {
			fmt.Println("An error occured while converting inputs to numbers", err)
		}

		list1 = append(list1, value1)
		list2 = append(list2, value2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	sum := 0

	for i := range list1 {
		sum += abs(list1[i] - list2[i])
	}

	fmt.Println(sum)
}

func abs(r int) int {
	if r < 0 {
		return -r
	}
	return r
}
