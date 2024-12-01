package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./part-2/input.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(data), "\r\n")
	var firsts []int
	var seconds []int

	for _, row := range rows {
		pair := strings.Split(row, "   ")
		first, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}

		firsts = append(firsts, first)
		seconds = append(seconds, second)
	}

	seconds_lookup := make(map[int]int)
	for _, second := range seconds {
		seconds_lookup[second] += 1
	}

	var total int

	for _, first := range firsts {
		if count, ok := seconds_lookup[first]; ok {
			total += count * first
		}
	}

	fmt.Println(total)
}
