package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type numbers []int

func readFile(filename string) numbers {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nmb := numbers{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		nmb = append(nmb, int(value))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nmb
}

func isValid(n int, preamle []int) bool {
	for i := 0; i < len(preamle)-1; i++ {
		for j := i + 1; j < len(preamle); j++ {
			if preamle[i]+preamle[j] == n {
				return true
			}
		}
	}
	return false
}

func getWrongNumber(nmb numbers, preambleSize int) int {
	preamble := []int{}
	for i := 0; i < preambleSize; i++ {
		preamble = append(preamble, nmb[i])
	}

	pIndex := 0
	for i := preambleSize; i < len(nmb); i++ {
		//fmt.Printf("preamble: %v\n", preamble)
		n := nmb[i]
		if !isValid(n, preamble) {
			return n
		}
		preamble[pIndex] = n
		pIndex = (pIndex + 1) % preambleSize
	}

	return 0
}

func task1(nmb numbers) int {
	return getWrongNumber(nmb, 25)
}

func getSet(wrongNumber int, nmb numbers) sort.IntSlice {
	set := sort.IntSlice{}

	for i := 0; i < len(nmb); i++ {
		sum := 0
		for j := i; j < len(nmb); j++ {
			sum += nmb[j]
			set = append(set, nmb[j])
			if sum == wrongNumber {
				return set
			}
			if sum > wrongNumber {
				break
			}
		}
		set = sort.IntSlice{}
		sum = 0
	}

	return sort.IntSlice{}
}

func getXMASNumber(wrongNumber int, nmb numbers) int {
	set := getSet(wrongNumber, nmb)

	sort.Sort(set)

	return set[0] + set[len(set)-1]
}

func task2(nmb numbers) int {
	wrongNumber := getWrongNumber(nmb, 25)
	return getXMASNumber(wrongNumber, nmb)
}

// task1: 776203571
// task2: 0

func main() {
	answers := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(answers))
	fmt.Printf("Task2: %v\n", task2(answers))
}
