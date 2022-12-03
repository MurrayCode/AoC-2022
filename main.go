package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day3"
)

func main() {
	content, err := os.ReadFile("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day3.Part2(string(content)))
}
