package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	var input = readFile(os.Args[1])

	var l, r []int

	for i := 0; i < len(input); i++ {
		pattern := regexp.MustCompile("^([0-9]+) {3}([0-9]+)")
		match := pattern.FindStringSubmatch(string(input[i]))

		if match != nil {
			i, err := strconv.Atoi(match[1])
			if err != nil {
				fmt.Println("Error parsing integer:", err)
				return
			}
			j, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println("Error parsing integer:", err)
				return
			}
			l = append(l, i)
			r = append(r, j)
		}
	}
	part1(l, r)
	part2(l, r)

}

func part1(l, r []int) {
	sort.Ints(l)
	sort.Ints(r)
	var diff int
	for i := 0; i < len(l); i++ {
		diff = diff + abs(l[i], r[i])
	}
	fmt.Printf("The solution for part 1 is %v\n", diff)
}

func part2(l, r []int) {
	var total int

	for i := 0; i < len(l); i++ {
		var q = 0
		for j := 0; j < len(r); j++ {
			if l[i] == r[j] {
				q += 1
			}
		}
		if q > 0 {
			total += l[i] * q
		}
	}
	fmt.Printf("The solution for part 2 is %v\n", total)
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func readFile(fileName string) []string {
	var lines []string

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
