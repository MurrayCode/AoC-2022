package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day2"
)

func main() {
	content, err := os.ReadFile("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(day2.Part2(string(content)))
}
