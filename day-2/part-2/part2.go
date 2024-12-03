package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func safe(levels []string) bool {
	sign := 0
	for i, level := range levels {
		if i == 0 {
			first, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(levels[1])
			if err != nil {
				panic(err)
			}
			if first > second {
				sign = -1
			} else if second > first {
				sign = 1
			} else {
				sign = 0
				break
			}

			continue
		}

		level, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}
		prev, err := strconv.Atoi(levels[i-1])
		if err != nil {
			panic(err)
		}
		delta := level - prev
		if delta*sign >= 1 && delta*sign <= 3 {
			continue
		} else {
			sign = 0
			break
		}
	}

	return sign != 0
}

func safeWithoutIndex(levels []string, index int) bool {
	sign := 0
	for i, level := range levels {
		if index == i {
			continue
		}
		if (i == 0 && index != 0) || (i == 1 && index == 0) {
			first, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			second, err := strconv.Atoi(levels[i+1])
			if err != nil {
				panic(err)
			}
			if i+1 == index {
				second, err = strconv.Atoi(levels[i+2])
				if err != nil {
					panic(err)
				}
			}
			if first > second {
				sign = -1
			} else if second > first {
				sign = 1
			} else {
				sign = 0
				break
			}
			continue
		}

		level, err := strconv.Atoi(level)
		if err != nil {
			panic(err)
		}

		prev, err := strconv.Atoi(levels[i-1])
		if err != nil {
			panic(err)
		}

		if i-1 == index {
			prev, err = strconv.Atoi(levels[i-2])
			if err != nil {
				panic(err)
			}
		}

		delta := level - prev

		if delta*sign < 1 || delta*sign > 3 {
			sign = 0
			break
		}

	}

	return sign != 0
}

func Run() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(file), "\r\n")

	count := 0

	for _, row := range rows {
		levels := strings.Split(row, " ")
		if safe(levels) {
			count += 1
			continue
		}
		for j := 0; j < len(levels); j++ {
			if safeWithoutIndex(levels, j) {
				count += 1
				break
			}
		}
	}
	fmt.Println(count)
}
