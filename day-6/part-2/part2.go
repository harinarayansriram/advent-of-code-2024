package part2

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type guard struct {
	i  int
	j  int
	di int
	dj int
}

func loops(old_obstacleIJ map[int][]int, old_obstacleJI map[int][]int, o_i int, o_j int, guard guard) bool {
	obstacleIJ := make(map[int][]int, len(old_obstacleIJ))
	obstacleJI := make(map[int][]int, len(old_obstacleJI))

	for i, obstacles := range old_obstacleIJ {
		obstacleIJ[i] = make([]int, len(obstacles))
		copy(obstacleIJ[i], obstacles)
	}
	for j, obstacles := range old_obstacleJI {
		obstacleJI[j] = make([]int, len(obstacles))
		copy(obstacleJI[j], obstacles)
	}

	obstacleIJ[o_i] = append(obstacleIJ[o_i], o_j)
	obstacleJI[o_j] = append(obstacleJI[o_j], o_i)

	slices.Sort(obstacleIJ[o_i])
	slices.Sort(obstacleJI[o_j])

	visited := make(map[string]bool)

	// turning: up -> right -> down -> left
	for {
		// up, so find largest value in obstacleJI[guard.j] that is less than guard.i
		if guard.di == -1 && guard.dj == 0 {
			col_obstacles := obstacleJI[guard.j]
			o_list_i, _ := slices.BinarySearch(col_obstacles, guard.i)
			if o_list_i == 0 {
				return false
			}

			o_i := col_obstacles[o_list_i-1]

			stateKey := fmt.Sprintf("%d,%d,%d,%d", o_i+1, guard.j, guard.di, guard.dj)
			if visited[stateKey] {
				return true
			}
			visited[stateKey] = true

			guard.i = o_i + 1
			// guard.j = guard.j
			guard.di = 0
			guard.dj = 1
			continue
		}

		// right, so find smallest value in obstacleIJ[guard.i] that is greater than guard.j
		if guard.di == 0 && guard.dj == 1 {
			row_obstacles := obstacleIJ[guard.i]
			o_list_j, _ := slices.BinarySearch(row_obstacles, guard.j)
			if o_list_j == len(row_obstacles) {
				return false
			}

			o_j := row_obstacles[o_list_j]

			stateKey := fmt.Sprintf("%d,%d,%d,%d", guard.i, o_j-1, guard.di, guard.dj)
			if visited[stateKey] {
				return true
			}
			visited[stateKey] = true

			// guard.i = guard.i
			guard.j = o_j - 1
			guard.di = 1
			guard.dj = 0
			continue
		}

		// down, so find smallest value in obstacleJI[guard.j] that is greater than guard.i
		if guard.di == 1 && guard.dj == 0 {
			col_obstacles := obstacleJI[guard.j]
			o_list_i, _ := slices.BinarySearch(col_obstacles, guard.i)
			if o_list_i == len(col_obstacles) {
				return false
			}

			o_i := col_obstacles[o_list_i]

			stateKey := fmt.Sprintf("%d,%d,%d,%d", o_i-1, guard.j, guard.di, guard.dj)
			if visited[stateKey] {
				return true
			}
			visited[stateKey] = true

			guard.i = o_i - 1
			// guard.j = guard.j
			guard.di = 0
			guard.dj = -1
			continue
		}

		// left, so find largest value in obstacleIJ[guard.i] that is less than guard.j
		if guard.di == 0 && guard.dj == -1 {
			row_obstacles := obstacleIJ[guard.i]
			o_list_j, _ := slices.BinarySearch(row_obstacles, guard.j)
			if o_list_j == 0 {
				return false
			}

			o_j := row_obstacles[o_list_j-1]

			stateKey := fmt.Sprintf("%d,%d,%d,%d", guard.i, o_j+1, guard.di, guard.dj)
			if visited[stateKey] {
				return true
			}
			visited[stateKey] = true

			// guard.i = guard.i
			guard.j = o_j + 1
			guard.di = -1
			guard.dj = 0
		}
	}
}
func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	obstacleIJ := make(map[int][]int, 0)
	obstacleJI := make(map[int][]int, 0)
	guard := guard{0, 0, -1, 0}

	rows := strings.Split(string(data), "\r\n")

	for i, row := range rows {
		for j, loc := range row {
			if loc == '#' {
				if obstacleIJ[i] == nil {
					obstacleIJ[i] = make([]int, 0)
				}
				if obstacleJI[j] == nil {
					obstacleJI[j] = make([]int, 0)
				}
				obstacleIJ[i] = append(obstacleIJ[i], j)
				obstacleJI[j] = append(obstacleJI[j], i)
			}
			if loc == '^' {
				guard.i = i
				guard.j = j
			}
		}
	}

	for _, obstacle_row := range obstacleIJ {
		slices.Sort(obstacle_row)
	}
	for _, obstacle_col := range obstacleJI {
		slices.Sort(obstacle_col)
	}

	total := 0
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[0]); j++ {
			if rows[i][j] != '#' && rows[i][j] != '^' && loops(obstacleIJ, obstacleJI, i, j, guard) {
				total += 1
			}
		}
	}

	fmt.Println(total)
}
