package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day4"
)

func main() {
	content, err := os.ReadFile("day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day4.Part2(string(content)))
}
