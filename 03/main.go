package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type treeRow []int
type treeMap []treeRow

func readFile(filename string) treeMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := treeMap{}
	scanner := bufio.NewScanner(file)
	lineLen := 0
	for scanner.Scan() {
		row := treeRow{}
		line := scanner.Text()
		if lineLen == 0 {
			lineLen = len(line)
		}
		if lineLen != len(line) {
			panic("Wrong line size")
		}

		for i := 0; i < lineLen; i++ {
			if string(line[i]) == "." {
				row = append(row, 0)
				continue
			}
			if string(line[i]) == "#" {
				row = append(row, 1)
				continue
			}
			panic("unexpected char: " + string(line[i]))
		}
		res = append(res, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	return res
}

func checkSlope(rightStep int, downStep int, tMap treeMap) int {
	rowNum := 0
	colNum := 0
	treeCount := 0
	rowLen := len(tMap[0])
	for {
		if tMap[rowNum][colNum] == 1 {
			treeCount++
		}
		if rowNum == len(tMap)-1 {
			break
		}
		rowNum += downStep
		colNum = (colNum + rightStep) % rowLen
	}
	return treeCount
}

func task1(tMap treeMap) int {
	return checkSlope(3, 1, tMap)
}

func task2(tMap treeMap) int {
	res := checkSlope(1, 1, tMap)
	res *= checkSlope(3, 1, tMap)
	res *= checkSlope(5, 1, tMap)
	res *= checkSlope(7, 1, tMap)
	res *= checkSlope(1, 2, tMap)
	return res
}

// task1: 234
// task2: 5813773056

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
