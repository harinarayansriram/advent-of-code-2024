package part1

import (
	"fmt"
	"os"
	"strings"
)

func countWith(i int, j int, di int, dj int, graph []string) int {
	acc := ""
	for i >= 0 && i < len(graph) && j >= 0 && j < len(graph[0]) {
		acc += string(graph[i][j])
		i += di
		j += dj
	}

	return strings.Count(acc, "XMAS") + strings.Count(acc, "SAMX")
}

func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	graph := strings.Split(string(data), "\r\n")
	total := 0
	for i := 0; i < len(graph); i++ {
		// rows
		total += countWith(i, 0, 0, 1, graph)
		// diagonal
		total += countWith(i, 0, 1, 1, graph)
		if i != len(graph)-1 {
			total += countWith(i, 0, -1, 1, graph)
		}
	}
	for j := 0; j < len(graph[0]); j++ {
		// columns
		total += countWith(0, j, 1, 0, graph)
		// diagonal
		if j != 0 {
			total += countWith(0, j, 1, 1, graph)
		}
		total += countWith(len(graph)-1, j, -1, 1, graph)
	}

	fmt.Println(total)
}
