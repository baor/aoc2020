package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	operation string
	argument  int
}

type program map[int]instruction

func createInstruction(input string) instruction {
	input = strings.TrimSpace(input)

	rgInstruction := regexp.MustCompile("(\\w+) ([-+]\\d+)")
	match := rgInstruction.FindStringSubmatch(input)
	if len(match) < 3 {
		panic("match<3! input: " + input)
	}

	value, err := strconv.ParseInt(match[2], 10, 32)
	if err != nil {
		panic(err)
	}

	instr := instruction{
		operation: match[1],
		argument:  int(value),
	}

	fmt.Printf("%+v\n", instr)
	return instr
}

func execute(prg program) (int, bool) {
	accumulator := 0
	i := 0
	visited := map[int]bool{}
	for {
		if visited[i] {
			fmt.Println("Loop is detected!")
			return accumulator, true
		}
		if i >= len(prg) {
			fmt.Println("Program ended!")
			return accumulator, false
		}

		visited[i] = true
		switch prg[i].operation {
		case "nop":
			i++
			continue
		case "acc":
			accumulator += prg[i].argument
			i++
			continue
		case "jmp":
			i += prg[i].argument
			continue
		default:
			panic("unknown operation :" + prg[i].operation)
		}
	}
}

func readFile(filename string) program {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	prg := program{}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		prg[i] = createInstruction(scanner.Text())
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return prg
}

func task1(prg program) int {
	accumulator, _ := execute(prg)
	return accumulator
}

func task2(prg program) int {
	for i := 0; i < len(prg); i++ {
		switch prg[i].operation {
		case "jmp":
			tempInstruction := prg[i]
			prg[i] = instruction{
				operation: "nop",
			}
			accumulator, hasLoop := execute(prg)
			if !hasLoop {
				return accumulator
			}
			prg[i] = tempInstruction
			continue
		case "nop":
			tempInstruction := prg[i]
			prg[i] = instruction{
				operation: "jmp",
				argument:  tempInstruction.argument,
			}
			accumulator, hasLoop := execute(prg)
			if !hasLoop {
				return accumulator
			}
			prg[i] = tempInstruction
			continue
		}
	}

	return 0
}

// task1: 1528
// task2: 34988

func main() {
	answers := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(answers))
	fmt.Printf("Task2: %v\n", task2(answers))
}
