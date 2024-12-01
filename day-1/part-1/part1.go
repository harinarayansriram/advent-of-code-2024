package part1

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./part-1/input.txt")
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

	slices.Sort(firsts)
	slices.Sort(seconds)

	var sum int
	for i := 0; i < len(firsts); i++ {
		diff := firsts[i] - seconds[i]
		if diff < 0 {
			diff *= -1
		}
		sum += diff
	}

	fmt.Println(sum)

}
