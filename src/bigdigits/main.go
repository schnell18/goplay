package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

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
		"  333 ",
		" 3   3",
		"     3",
		"   33 ",
		"     3",
		" 3   3",
		"  333 ",
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
		"     6 ",
		"    6  ",
		"   6   ",
		"  6  6 ",
		" 6    6",
		" 6    6",
		"  6666 ",
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
		" 88888 ",
		"8     8",
		" 8   8 ",
		"   8   ",
		" 8   8 ",
		"8     8",
		" 88888 ",
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

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	digits := os.Args[1]
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
		fmt.Println(line)
	}
}
