package main

import (
	"fmt"
	"os"
	"utils"
)

const XMAS = "XMAS"

func main() {
	var lines = utils.ReadFile(os.Args[1])
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {

			total += checkXMAS("north", lines, i, j, 0)     //2
			total += checkXMAS("south", lines, i, j, 0)     //1
			total += checkXMAS("east", lines, i, j, 0)      //3
			total += checkXMAS("west", lines, i, j, 0)      //2
			total += checkXMAS("northeast", lines, i, j, 0) //4
			total += checkXMAS("southeast", lines, i, j, 0) //1
			total += checkXMAS("northwest", lines, i, j, 0) //4
			total += checkXMAS("southwest", lines, i, j, 0) //1
		}
	}
	fmt.Printf("The total XMAS: %v\n", total)
}

func part2(lines []string) {
	total := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j] == 'A' && i+1 <= len(lines)-1 && i-1 >= 0 && j+1 <= len(lines)-1 && j-1 >= 0 {
				if lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' && lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S' {
					total += 1
				}
				if lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' && lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M' {
					total += 1
				}
				if lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M' && lines[i+1][j-1] == 'S' && lines[i-1][j+1] == 'M' {
					total += 1
				}
				if lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M' && lines[i+1][j-1] == 'M' && lines[i-1][j+1] == 'S' {
					total += 1
				}
			}
		}
	}
	fmt.Printf("The total X-MAS: %v\n", total)
}

func checkXMAS(direction string, lines []string, i, j, pos int) int {
	//fmt.Printf("i:%v j:%v val:%v | pos:%v const:%v\n", i, j, string(lines[i][j]), pos, string(XMAS[pos]))
	if lines[i][j] == XMAS[pos] && pos == 3 {
		return 1
	}
	switch direction {
	case "north":
		{
			if lines[i][j] == XMAS[pos] && i-1 >= 0 {
				return checkXMAS(direction, lines, i-1, j, pos+1)
			}
		}
	case "south":
		{
			if lines[i][j] == XMAS[pos] && i+1 <= len(lines)-1 {
				return checkXMAS(direction, lines, i+1, j, pos+1)
			}
		}
	case "east":
		{
			if lines[i][j] == XMAS[pos] && j+1 <= len(lines)-1 {
				return checkXMAS(direction, lines, i, j+1, pos+1)
			}
		}
	case "west":
		{
			if lines[i][j] == XMAS[pos] && j-1 >= 0 {
				return checkXMAS(direction, lines, i, j-1, pos+1)
			}
		}
	case "northeast":
		{
			if lines[i][j] == XMAS[pos] && i-1 >= 0 && j+1 <= len(lines)-1 {
				return checkXMAS(direction, lines, i-1, j+1, pos+1)
			}
		}
	case "southeast":
		{
			if lines[i][j] == XMAS[pos] && i+1 <= len(lines)-1 && j+1 <= len(lines)-1 {
				return checkXMAS(direction, lines, i+1, j+1, pos+1)
			}
		}
	case "northwest":
		{
			if lines[i][j] == XMAS[pos] && i-1 >= 0 && j-1 >= 0 {
				return checkXMAS(direction, lines, i-1, j-1, pos+1)
			}
		}
	case "southwest":
		{
			if lines[i][j] == XMAS[pos] && i+1 <= len(lines)-1 && j-1 >= 0 {
				return checkXMAS(direction, lines, i+1, j-1, pos+1)
			}
		}

	}
	return 0
}
