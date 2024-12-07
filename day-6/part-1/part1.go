package part1

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

func fillFrom(i_i, j_i, i_f, j_f int, locations [][]bool) {
	di := 0
	if i_f > i_i {
		di = 1
	}
	if i_f < i_i {
		di = -1
	}
	dj := 0
	if j_f > j_i {
		dj = 1
	}
	if j_f < j_i {
		dj = -1
	}
	i := i_i
	j := j_i
	for i != i_f || j != j_f {
		locations[i][j] = true
		i += di
		j += dj
	}
	locations[i][j] = true
}

func countLocations(locations [][]bool) int {
	count := 0
	for _, row := range locations {
		for _, loc := range row {
			if loc {
				count += 1
			}
		}
	}
	return count
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
	locations := make([][]bool, len(rows))

	for i, row := range rows {
		locations[i] = make([]bool, len(row))
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
	// turning: up -> right -> down -> left
	for {
		locations[guard.i][guard.j] = true
		// up, so find largest value in obstacleJI[guard.j] that is less than guard.i
		if guard.di == -1 && guard.dj == 0 {
			col_obstacles := obstacleJI[guard.j]
			o_list_i, _ := slices.BinarySearch(col_obstacles, guard.i)
			if o_list_i == 0 {
				// fill to the start of the column
				fillFrom(guard.i, guard.j, 0, guard.j, locations)
				break
			}

			o_i := col_obstacles[o_list_i-1]
			fillFrom(guard.i, guard.j, o_i+1, guard.j, locations)

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
				// fill to the end of the row
				fillFrom(guard.i, guard.j, guard.i, len(locations[0])-1, locations)
				break
			}

			o_j := row_obstacles[o_list_j]
			fillFrom(guard.i, guard.j, guard.i, o_j-1, locations)

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
				// fill to the end of the column
				fillFrom(guard.i, guard.j, len(locations)-1, guard.j, locations)
				break
			}

			o_i := col_obstacles[o_list_i]
			fillFrom(guard.i, guard.j, o_i-1, guard.j, locations)

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
				// fill to the start of the row
				fillFrom(guard.i, guard.j, guard.i, 0, locations)
				break
			}

			o_j := row_obstacles[o_list_j-1]
			fillFrom(guard.i, guard.j, guard.i, o_j+1, locations)

			// guard.i = guard.i
			guard.j = o_j + 1
			guard.di = -1
			guard.dj = 0
			continue
		}
	}

	fmt.Println(countLocations(locations))
}
