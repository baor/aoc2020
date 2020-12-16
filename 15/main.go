package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task1(input string, target int) int {
	substr := strings.Split(input, ",")
	numbs := []int{}
	for _, s := range substr {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			panic(err)
		}
		numbs = append(numbs, int(v))
	}
	fmt.Printf("numbs :%v\n", numbs)

	lastHeard := map[int][]int{}
	for i := 0; i < len(numbs); i++ {
		lastHeard[numbs[i]] = []int{i + 1}
	}

	iteration := len(numbs)
	prevSaid := numbs[len(numbs)-1]
	for {

		if len(lastHeard[prevSaid]) == 0 {
			lastHeard[prevSaid] = []int{iteration}
		} else {
			lastHeard[prevSaid] = []int{iteration, lastHeard[prevSaid][0]}
		}

		iteration++
		newSaid := 0
		if len(lastHeard[prevSaid]) == 2 {
			newSaid = lastHeard[prevSaid][0] - lastHeard[prevSaid][1]
		}

		//fmt.Printf("[%v] prev said %v, last time heard %v, so new %v\n", iteration, prevSaid, lastTimeHeard, newSaid)

		if iteration == target {
			return newSaid
		}
		prevSaid = newSaid
	}
}
func task2(input string, target int) int {
	return 0
}

// task1: 468
// task2: 0

func main() {
	input := "6,19,0,5,7,13,1"
	fmt.Printf("Task1: %v\n", task1(input, 2020))
	fmt.Printf("Task2: %v\n", task1(input, 30000000))
}
