package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"utils"
)

func main() {
	var lines = utils.ReadFile(os.Args[1])
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		//fmt.Printf("%v\n", lines[i])

		numbers := strings.Split(strings.TrimSpace(lines[i]), " ")

		if isSafe(numbers) {
			total++
		}
		//fmt.Printf("Dec: %v | Inc: %v\n", dec, inc)
		/*fmt.Printf("Numbers: %v | Total: %v\n", lines[i], total)*/
	}
	fmt.Printf("%v reports are safe\n", total)
}

func part2(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		//fmt.Printf("%v\n", lines[i])

		numbers := strings.Split(strings.TrimSpace(lines[i]), " ")

		if isSafe(numbers) || isSafeWithDampener(numbers) {
			total++
		}
		//fmt.Printf("Dec: %v | Inc: %v\n", dec, inc)
		/*fmt.Printf("Numbers: %v | Total: %v\n", lines[i], total)*/
	}
	fmt.Printf("Thanks to the Problem Dampener, %v reports are actually safe!\n", total)
}

func isSafeWithDampener(numbers []string) bool {
	for i := 0; i < len(numbers); i++ {
		var l = append([]string{}, numbers[:i]...)
		var r = append([]string{}, numbers[i+1:]...)
		var n []string
		n = append(n, l...)
		n = append(n, r...)
		if isSafe(n) {
			return true
		}
	}
	return false
}

func isSafe(numbers []string) bool {
	var prev int
	var inc = true
	var dec = true

	for i := 0; i < len(numbers); i++ {
		n, err := strconv.Atoi(strings.TrimSpace(numbers[i]))
		if err != nil {
			log.Fatal(err)
		}

		if i != 0 {
			diff := prev - n
			dec = abs(diff) <= 3 && abs(diff) >= 1 && diff > 0 && dec
			inc = abs(diff) <= 3 && abs(diff) >= 1 && diff < 0 && inc

		}
		prev = n
	}
	return inc != dec
}
func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
