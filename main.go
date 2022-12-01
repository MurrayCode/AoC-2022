package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day1"
)

func main() {
	content, err := os.ReadFile("day1/part1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(day1.Part2(string(content)))
}
