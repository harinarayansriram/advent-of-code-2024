package part1

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func evalute(expression string) int {
	stripped := expression[4 : len(expression)-1]
	pair := strings.Split(stripped, ",")
	first, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(err)
	}

	second, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(err)
	}
	return first * second
}
func Run() {
	data, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(`mul\([0-9]{1,3},[0-9]{1,3}\)`)
	matches := r.FindAllString(string(data), -1)
	sum := 0
	for _, match := range matches {
		sum += evalute(match)
	}
	fmt.Println(sum)
}
