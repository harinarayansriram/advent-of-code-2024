package part2

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	rows := strings.Split(string(data), "\r\n")

	positions := make(map[rune][][2]int)

	for i, row := range rows {
		for j, c := range row {
			if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				if positions[c] == nil {
					positions[c] = make([][2]int, 0)
				}
				positions[c] = append(positions[c], [2]int{i, j})
			}
		}
	}

	antinodes := make([][]int, len(rows))

	for i := range antinodes {
		antinodes[i] = make([]int, len(rows[0]))
	}

	for _, locs := range positions {
		for i, loc1 := range locs {
			for j := i + 1; j < len(locs); j++ {
				loc2 := locs[j]

				di := loc2[0] - loc1[0]
				dj := loc2[1] - loc1[1]

				k := 0
				for loc1[0]-di*k >= 0 && loc1[0]-di*k < len(rows) && loc1[1]-dj*k >= 0 && loc1[1]-dj*k < len(rows[0]) {
					antinodes[loc1[0]-di*k][loc1[1]-dj*k] = 1
					k += 1
				}

				k = 0
				for loc2[0]+di*k >= 0 && loc2[0]+di*k < len(rows) && loc2[1]+dj*k >= 0 && loc2[1]+dj*k < len(rows[0]) {
					antinodes[loc2[0]+di*k][loc2[1]+dj*k] = 1
					k += 1
				}
			}
		}
	}

	count := 0

	for _, row := range antinodes {
		for _, loc := range row {
			count += loc
		}
	}

	fmt.Println(count)
}
