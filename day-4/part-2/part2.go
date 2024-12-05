package part2

import (
	"fmt"
	"os"
	"strings"
)

func checkXMas(x int, y int, graph []string, chars string) bool {
	return graph[x][y] == chars[0] && graph[x][y+2] == chars[1] && graph[x+1][y+1] == chars[2] && graph[x+2][y] == chars[3] && graph[x+2][y+2] == chars[4]
}

func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	graph := strings.Split(string(data), "\r\n")
	// 4 possibilities
	// M.S
	// .A.
	// M.S

	// S.S
	// .A.
	// M.M

	// M.M
	// .A.
	// S.S

	// S.M
	// .A.
	// S.M
	count := 0
	fmt.Println("part 2", len(data))
	for i := 0; i < len(graph)-2; i++ {
		for j := 0; j < len(graph[0])-2; j++ {
			if checkXMas(i, j, graph, "MSAMS") || checkXMas(i, j, graph, "SSAMM") || checkXMas(i, j, graph, "MMASS") || checkXMas(i, j, graph, "SMASM") {
				count += 1
			}
		}

	}

	fmt.Println(count)
}
