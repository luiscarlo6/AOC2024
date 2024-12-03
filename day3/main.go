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
	part2(lines)
}

func part2(lines []string) {
	total := 0

	active := true
	for i := 0; i < len(lines); i++ {
		pattern := regexp.MustCompile("(mul)\\(([0-9]{1,3}),([0-9]{1,3})\\)|(do)\\(\\)|(don't)\\(\\)")
		match := pattern.FindAllStringSubmatch(lines[i], -1)

		if match != nil {
			for j := 0; j < len(match); j++ {
				//fmt.Printf("%v %v\n", match[j], active)

				var command = "mul"
				if match[j][1] == "mul" {
					command = match[j][1]
				} else if match[j][4] == "do" {
					command = match[j][4]
				} else if match[j][5] == "don't" {
					command = match[j][5]
				} else {
					fmt.Println("Invalid command")
				}

				switch command {
				case "mul":
					if active {
						a, _ := strconv.Atoi(match[j][2])
						b, _ := strconv.Atoi(match[j][3])
						total += a * b
					}
				case "do":
					active = true
				case "don't":
					active = false
				default:
					fmt.Println("Invalid command")
				}

			}
		}
	}

	fmt.Printf("The result of the enabled/disabled multiplications is: %v\n", total)
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
