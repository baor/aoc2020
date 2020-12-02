package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type lineWithPassword struct {
	minNum   int
	maxNum   int
	char     string
	password string
}

func readFile(filename string) []lineWithPassword {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	res := []lineWithPassword{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var minNum, maxNum int
		var char, password string
		// 1-3 a: abcde
		fmt.Sscanf(scanner.Text(), "%d-%d %s %s", &minNum, &maxNum, &char, &password)
		char = strings.Trim(char, ":")
		res = append(res, lineWithPassword{minNum: minNum, maxNum: maxNum, char: char, password: password})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", res)
	return res
}

func task1IsValid(line lineWithPassword) bool {
	entriesCount := strings.Count(line.password, line.char)
	return entriesCount >= line.minNum && entriesCount <= line.maxNum
}

func task1(input []lineWithPassword) int {
	numberOfValidPasswords := 0

	for _, line := range input {
		if len(line.char) == 0 {
			panic("empty char")
		}
		if task1IsValid(line) {
			numberOfValidPasswords++
		}
	}

	return numberOfValidPasswords
}

func task2IsValid(line lineWithPassword) bool {
	if line.password[line.minNum-1] == line.char[0] && line.password[line.maxNum-1] != line.char[0] {
		return true
	}
	if line.password[line.minNum-1] != line.char[0] && line.password[line.maxNum-1] == line.char[0] {
		return true
	}
	return false
}

func task2(input []lineWithPassword) int {
	numberOfValidPasswords := 0

	for _, line := range input {
		if len(line.char) == 0 {
			panic("empty char")
		}
		if task2IsValid(line) {
			numberOfValidPasswords++
		}
	}

	return numberOfValidPasswords
}

// task1: 607
// task2: 321

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
