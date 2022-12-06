package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MurrayCode/AoC-2022/day6"
)

func main() {
	content, err := os.ReadFile("day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(day6.Part2(string(content)))
}
