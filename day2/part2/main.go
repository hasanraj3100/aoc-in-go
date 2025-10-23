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

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		var list []int

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting input to a number")
				continue
			}

			list = append(list, num)
		}

		if checkValidity(list) || canMakeValid(list) {
			sum++
		}

	}

	fmt.Println(sum)
}

func checkValidity(list []int) bool {
	inc := list[1] > list[0]

	for i := 1; i < len(list); i++ {
		if inc && list[i] <= list[i-1] {
			return false
		}
		if !inc && list[i] >= list[i-1] {
			return false
		}
		if abs(list[i]-list[i-1]) < 1 || abs(list[i]-list[i-1]) > 3 {
			return false
		}

	}

	return true
}

func canMakeValid(list []int) bool {
	for i := range list {
		temp := make([]int, 0)
		temp = append(temp, list[:i]...)
		temp = append(temp, list[i+1:]...)
		if checkValidity(temp) {
			return true
		}
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
