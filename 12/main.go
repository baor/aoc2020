package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	action string
	value  int
}

type ship struct {
	x         int
	y         int
	direction string
}

var rotationRightSeq = "ESWN"

func (s *ship) move(direction string, value int) {
	switch direction {
	case "E":
		s.x += value
	case "N":
		s.y += value
	case "W":
		s.x -= value
	case "S":
		s.y -= value
	default:
		panic("move error!")
	}
}

func (s *ship) rotateLeft(value int) {
	numberOfPositionChanges := value / 90
	shift := numberOfPositionChanges % len(rotationRightSeq)

	currentDirectionIndex := strings.Index(rotationRightSeq, s.direction)
	newDirectionIndex := currentDirectionIndex - shift
	if newDirectionIndex < 0 {
		newDirectionIndex += len(rotationRightSeq)
	}
	s.direction = string(rotationRightSeq[newDirectionIndex])
}

func (s *ship) rotateRight(value int) {
	numberOfPositionChanges := value / 90
	shift := numberOfPositionChanges % len(rotationRightSeq)

	currentDirectionIndex := strings.Index(rotationRightSeq, s.direction)
	newDirectionIndex := currentDirectionIndex + shift
	if newDirectionIndex >= len(rotationRightSeq) {
		newDirectionIndex -= len(rotationRightSeq)
	}
	s.direction = string(rotationRightSeq[newDirectionIndex])
}

func (s *ship) moveForward(value int) {
	s.move(s.direction, value)
}

func (s *ship) moveToTheWaypoint(value int, wp waypoint) {
	s.x += wp.xRel * value
	s.y += wp.yRel * value
}

type waypoint struct {
	xRel int
	yRel int
}

func (wp *waypoint) move(direction string, value int) {
	switch direction {
	case "E":
		wp.xRel += value
	case "N":
		wp.yRel += value
	case "W":
		wp.xRel -= value
	case "S":
		wp.yRel -= value
	default:
		panic("move error!")
	}
}

func (wp *waypoint) rotateLeft(value int) {
	numberOfPositionChanges := value / 90

	for i := 0; i < numberOfPositionChanges; i++ {
		temp := wp.xRel
		wp.xRel = -1 * wp.yRel
		wp.yRel = temp
	}
}

func (wp *waypoint) rotateRight(value int) {
	numberOfPositionChanges := value / 90

	for i := 0; i < numberOfPositionChanges; i++ {
		temp := wp.xRel
		wp.xRel = wp.yRel
		wp.yRel = -1 * temp
	}
}

func createInstruction(input string) instruction {
	instr := instruction{
		action: string(input[0]),
	}

	val, err := strconv.ParseInt(input[1:], 10, 64)
	if err != nil {
		panic(err)
	}
	instr.value = int(val)

	return instr
}

func readFile(filename string) []instruction {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	instructions := []instruction{}
	for scanner.Scan() {
		instructions = append(instructions, createInstruction(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", instructions)
	return instructions
}

func doInstructionsTask1(instructions []instruction) ship {
	s := ship{
		x:         0,
		y:         0,
		direction: "E",
	}
	for _, instr := range instructions {
		switch instr.action {
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			s.move(instr.action, instr.value)
		case "L":
			s.rotateLeft(instr.value)
		case "R":
			s.rotateRight(instr.value)
		case "F":
			s.moveForward(instr.value)
		default:
			panic("wrong input!")
		}
	}
	return s
}

func doInstructionsTask2(instructions []instruction) ship {
	s := ship{
		x:         0,
		y:         0,
		direction: "E",
	}

	wp := waypoint{
		xRel: 10,
		yRel: 1,
	}

	for _, instr := range instructions {
		switch instr.action {
		case "N":
			fallthrough
		case "S":
			fallthrough
		case "E":
			fallthrough
		case "W":
			wp.move(instr.action, instr.value)
		case "L":
			wp.rotateLeft(instr.value)
		case "R":
			wp.rotateRight(instr.value)
		case "F":
			s.moveToTheWaypoint(instr.value, wp)
		default:
			panic("wrong input!")
		}
	}
	return s
}

func task1(instructions []instruction) int {
	s := doInstructionsTask1(instructions)
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func task2(instructions []instruction) int {
	s := doInstructionsTask2(instructions)
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

// task1: 1838
// task2: 89936

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
