package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day5"
)

func main() {
	content, err := os.ReadFile("day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day5.Part1(string(content)))
}
