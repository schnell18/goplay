package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	width := 0
	printBar := false
	digits := os.Args[1]
	if len(os.Args) == 1 {
		usage()
		os.Exit(1)
	} else if len(os.Args) == 3 {
		if os.Args[1] == "-b" || os.Args[1] == "--bar" {
			printBar = true
			digits = os.Args[2]
		} else {
			usage()
			os.Exit(1)
		}
	} else if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		usage()
		os.Exit(0)
	}

	for row := range bigDigits[0] {
		line := ""

		for column := range digits {
			digit := digits[column] - '0'
			if digit >= 0 && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("Invalid whole number")
			}
		}
		width = len(line)
		if row == 0 && printBar {
			fmt.Println(strings.Repeat("*", width))
		}
		fmt.Println(line)
	}
	if printBar {
		fmt.Println(strings.Repeat("*", width))
	}
}

func usage() {
	fmt.Printf("Usage: %s [-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
	fmt.Println("-b --bar   draw an underbar and an overbar")
	fmt.Println("-h --help  show usage and exit")
}

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},

	{
		" 1 ",
		"11 ",
		" 1 ",
		" 1 ",
		" 1 ",
		" 1 ",
		"111",
	},

	{
		" 222 ",
		"2   2",
		"   2 ",
		"  2  ",
		" 2   ",
		"2    ",
		"22222",
	},

	{
		" 333 ",
		"3   3",
		"    3",
		"  33 ",
		"    3",
		"3   3",
		" 333 ",
	},

	{
		"    4  ",
		"   44  ",
		"  4 4  ",
		" 4  4  ",
		" 444444",
		"    4  ",
		"    4  ",
	},

	{
		"55555",
		"5    ",
		"5    ",
		" 555 ",
		"    5",
		"5   5",
		" 555 ",
	},

	{
		"  666 ",
		" 6    ",
		" 6    ",
		" 6666 ",
		" 6   6",
		" 6   6",
		"  666 ",
	},

	{
		"77777",
		"    7",
		"   7 ",
		"  7  ",
		" 7   ",
		"7    ",
		"7    ",
	},

	{
		" 888 ",
		"8   8",
		"8   8",
		" 888 ",
		"8   8",
		"8   8",
		" 888 ",
	},

	{
		" 9999",
		"9   9",
		"9   9",
		" 9999",
		"    9",
		"   9 ",
		"  9  ",
	},
}
