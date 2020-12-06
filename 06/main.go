package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type group struct {
	numPeople int
	answers   answers
}

type answers map[string]int

func createAnswersOfTheGroup(input string) answers {
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\n", "")
	input = strings.Join(strings.Fields(input), "")
	match, err := regexp.MatchString("^[a-z]+$", input)
	if !match || err != nil {
		panic("str err: \"" + input + "\"")
	}
	ans := answers{}
	for i := 0; i < len(input); i++ {
		char := string(input[i])
		ans[char]++
	}
	return ans
}

func readFile(filename string) []group {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	groups := []group{}
	scanner := bufio.NewScanner(file)
	inputForOneGroup := ""
	numPeople := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			groups = append(groups, group{
				numPeople: numPeople,
				answers:   createAnswersOfTheGroup(inputForOneGroup),
			})
			inputForOneGroup = ""
			numPeople = 0
			continue
		}
		inputForOneGroup += scanner.Text()
		numPeople++
	}

	if inputForOneGroup != "" {
		groups = append(groups, group{
			numPeople: numPeople,
			answers:   createAnswersOfTheGroup(inputForOneGroup),
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", passports)
	return groups
}

func task1(groups []group) int {
	counts := 0
	for _, g := range groups {
		counts += len(g.answers)
	}

	return counts
}

func task2(groups []group) int {
	counts := 0
	//fmt.Printf("Task2 groups: %v\n", groups)
	for _, g := range groups {
		for _, ansCount := range g.answers {
			if g.numPeople == ansCount {
				counts++
			}
		}
	}
	return counts
}

// task1: 6521
// task2: 3305

func main() {
	answers := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(answers))
	fmt.Printf("Task2: %v\n", task2(answers))
}
