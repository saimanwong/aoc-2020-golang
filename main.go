package main

import (
	"fmt"
	"os"
	"regexp"

	f "./utils/functions"

	"./problems/day01"
	"./problems/day02"
	"./problems/day03"
	"./problems/day04"
	"./problems/day05"
	"./problems/day06"
	"./problems/day07"
	"./problems/day08"
	"./problems/day09"
	"./problems/day10"
	"./problems/day11"
	"./problems/day12"
	"./problems/day13"
	"./problems/day14"
	"./problems/day15"
	"./problems/day16"
	"./problems/day17"
	"./problems/day18"
	"./problems/day19"
	"./problems/day20"
	"./problems/day21"
	"./problems/day22"
	"./problems/day23"
	"./problems/day24"
	"./problems/day25"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("exactly 2 arguments required")
		os.Exit(1)
	}

	day := os.Args[1]
	re := regexp.MustCompile("^[0-9]{2}$")
	if !re.MatchString(day) {
		fmt.Println("first argument must only contain numbers")
		os.Exit(2)
	}

	inputFile := os.Args[2]
	re = regexp.MustCompile("^[a-z0-9]+$")
	if !re.MatchString(inputFile) {
		fmt.Println("second argument must only contain alphanumeric characters")
		os.Exit(3)
	}

	path := "problems/day" + day + "/inputs/" + inputFile
	input := f.ReadFile(path)

	switch day {
	case "01":
		day01.Run(input)
	case "02":
		day02.Run(input)
	case "03":
		day03.Run(input)
	case "04":
		day04.Run(input)
	case "05":
		day05.Run(input)
	case "06":
		day06.Run(input)
	case "07":
		day07.Run(input)
	case "08":
		day08.Run(input)
	case "09":
		day09.Run(input)
	case "10":
		day10.Run(input)
	case "11":
		day11.Run(input)
	case "12":
		day12.Run(input)
	case "13":
		day13.Run(input)
	case "14":
		day14.Run(input)
	case "15":
		day15.Run(input)
	case "16":
		day16.Run(input)
	case "17":
		day17.Run(input)
	case "18":
		day18.Run(input)
	case "19":
		day19.Run(input)
	case "20":
		day20.Run(input)
	case "21":
		day21.Run(input)
	case "22":
		day22.Run(input)
	case "23":
		day23.Run(input)
	case "24":
		day24.Run(input)
	case "25":
		day25.Run(input)
	default:
		fmt.Printf("%s not implemented yet...\n", day)
	}
}
