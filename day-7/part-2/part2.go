package part2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func concat(a, b int) int {
	n, err := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	if err != nil {
		panic(err)
	}
	return n
}
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
		
		if concat(current[0], equation[current[1]+1]) <= val {
			stack = append(stack, [2]int{concat(current[0], equation[current[1]+1]), current[1] + 1})
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

//	func dfs(val int, equation []int, current [2]int) bool {
//		if current[1] == len(equation)-1 {
//			return current[0] == val
//		}
//		if current[0]*equation[current[1]+1] <= val && dfs(val, equation, [2]int{current[0] * equation[current[1]+1], current[1] + 1}) {
//			return true
//		}
//		if current[0]+equation[current[1]+1] <= val && dfs(val, equation, [2]int{current[0] + equation[current[1]+1], current[1] + 1}) {
//			return true
//		}
//		return false
//	}
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

	// fmt.Println(concat(12, 1345))

	fmt.Println(total)
}
