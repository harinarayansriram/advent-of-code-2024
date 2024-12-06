package part2

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func toposort(inEdgeCounts map[int]int, forwardEdges map[int][]int) []int {
	stack := make([]int, 0)
	ordering := make([]int, 0)

	for k, v := range inEdgeCounts {
		if v == 0 {
			stack = append(stack, k)
		}
	}

	for len(stack) > 0 {
		// pop off the stack
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ordering = append(ordering, current)

		for _, next := range forwardEdges[current] {
			inEdgeCounts[next] -= 1
			if inEdgeCounts[next] == 0 {
				stack = append(stack, next)
			}
		}
	}

	return ordering
}

func forwardEdgesWithNodes(forwardEdges map[int][]int, nodes []int) map[int][]int {
	newForwardEdges := make(map[int][]int)
	node_set := make(map[int]struct{})
	for _, node := range nodes {
		node_set[node] = struct{}{}
	}
	for _, node := range nodes {
		newForwardEdges[node] = make([]int, 0)
		for _, edge := range forwardEdges[node] {
			if _, ok := node_set[edge]; ok {
				newForwardEdges[node] = append(newForwardEdges[node], edge)
			}
		}
	}
	return newForwardEdges
}

func getInEdgeCounts(forwardEdges map[int][]int) map[int]int {
	inEdgeCounts := make(map[int]int)
	for i, edges := range forwardEdges {
		if _, ok := inEdgeCounts[i]; !ok {
			inEdgeCounts[i] = 0
		}
		for _, edge := range edges {
			inEdgeCounts[edge] += 1
		}
	}
	return inEdgeCounts
}
func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	forwardEdges := make(map[int][]int)

	rows := strings.Split(string(data), "\r\n")

	nextSectionStart := 0

	for i, row := range rows {
		if row == "" {
			nextSectionStart = i + 1
			break
		}
		pair := strings.Split(row, "|")
		first, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}
		var edges []int
		var ok bool

		if edges, ok = forwardEdges[first]; !ok {
			forwardEdges[first] = make([]int, 0)
		}

		forwardEdges[first] = append(edges, second)
	}

	total := 0

	for i := nextSectionStart; i < len(rows); i++ {
		row := strings.Split(rows[i], ",")
		updates := make([]int, 0)

		for _, v := range row {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			updates = append(updates, num)
		}

		limited_fwd_edges := forwardEdgesWithNodes(forwardEdges, updates)
		sorted := toposort(getInEdgeCounts(limited_fwd_edges), limited_fwd_edges)

		sorted_lookup := make(map[int]int)
		for i, v := range sorted {
			sorted_lookup[v] = i
		}

		ordered := slices.IsSortedFunc(updates, func(a, b int) int {
			return sorted_lookup[a] - sorted_lookup[b]
		})
		if !ordered {
			total += sorted[len(sorted)/2]
		}
	}

	fmt.Println(total)

}
