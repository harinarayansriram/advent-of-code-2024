package part1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(val int, equation []int) bool {
	stack := make([][2]int, 0)
	stack = append(stack, [2]int{equation[0], 0})

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current[1] == len(equation)-1 {
			if current[0] == val {
				return true
			}

			continue
		}

		if current[0]*equation[current[1]+1] <= val {
			stack = append(stack, [2]int{current[0] * equation[current[1]+1], current[1] + 1})
		}

		if current[0]+equation[current[1]+1] <= val {
			stack = append(stack, [2]int{current[0] + equation[current[1]+1], current[1] + 1})
		}
		
	}
	return false
}

func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(data), "\r\n")
	valuesAndEquations := make(map[int][]int)
	for _, row := range rows {
		numAndVals := strings.Split(row, ": ")
		num, err := strconv.Atoi(numAndVals[0])
		if err != nil {
			panic(err)
		}
		equation := make([]int, 0)
		vals := strings.Split(numAndVals[1], " ")
		for _, val := range vals {
			parsedVal, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			equation = append(equation, parsedVal)
		}

		valuesAndEquations[num] = equation
	}
	total := 0
	for val, equation := range valuesAndEquations {
		if dfs(val, equation) {
			total += val
		}
	}

	fmt.Println(total)
}
