package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"utils"
)

func main() {
	lines := utils.ReadFile(os.Args[1])
	part1(lines)
}

func part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		pattern := regexp.MustCompile("mul\\(([0-9]{1,3}),([0-9]{1,3})\\)")
		match := pattern.FindAllStringSubmatch(lines[i], -1)

		if match != nil {
			for j := 0; j < len(match); j++ {
				a, _ := strconv.Atoi(match[j][1])
				b, _ := strconv.Atoi(match[j][2])
				total += a * b
			}
		}
	}

	fmt.Printf("The result of the multiplications is: %v\n", total)
}
