package main

import (
	"fmt"
	"reflect"
)

func main() {
	startPoint := []int{0, 0}
	start := [][]int{startPoint}
	fmt.Println(move(start))
}

func move(log [][]int) int {
	N := 12

	if len(log) == N+1 {
		return 1
	}

	cnt := 0
	north := []int{0, 1}
	south := []int{0, -1}
	west := []int{1, 0}
	east := []int{-1, 0}
	direction := [][]int{north, south, west, east}

	for _, d := range direction {
		nextPos := []int{log[len(log)-1][0] + d[0], log[len(log)-1][1] + d[1]}
		if !include(log, nextPos) {
			cnt += move(append(log, nextPos))
		}
	}

	return cnt
}

func include(log [][]int, nextPos []int) bool {
	for _, v := range log {
		if reflect.DeepEqual(v, nextPos) {
			return true
		}
	}
	return false
}
