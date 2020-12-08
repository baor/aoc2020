package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var a int
		fmt.Sscanf(scanner.Text(), "%d", &a)
		res = append(res, a)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return res
}

func task1(input []int) int {
	for i := 0; i < len(input)-1; i++ {
		for j := i; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}
	return 0
}

func task2(input []int) int {
	for i := 0; i < len(input)-2; i++ {
		for j := i; j < len(input)-1; j++ {
			for k := j; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}
	return 0
}

// task1: 970816
// task2: 96047280

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
