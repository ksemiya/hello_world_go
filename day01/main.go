package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Println("Part 1: ", solvePart1(lines))
	fmt.Println("Part 2: ", solvePart2(lines))
}

func solvePart1(lines []string) int {
	currState := 50
	ans := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		currSign, currNum := parseLine(line)
		currState += currNum * currSign
		currState %= 100
		if currState == 0 {
			ans++
		}
	}
	return ans
}

func solvePart2(lines []string) int {
	currState := 50
	ans := 0
	for _, line := range lines {
		if line == "" {
			break
		}
		currSign, currNum := parseLine(line)
		currZero := currState == 0
		currState += currNum * currSign
		if (currState < 0) && !currZero {
			ans++
		}
		ans += max(currState, -currState) / 100
		if currState == 0 {
			ans++
		}
		currState %= 100
		if currState < 0 {
			currState += 100
		}
	}
	return ans
}

func parseLine(line string) (sign int, num int) {
	sign = 1
	if line[0] == 'L' {
		sign = -1
	}
	var err error
	num, err = strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}
	return sign, num
}
