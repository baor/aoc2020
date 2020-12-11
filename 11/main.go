package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type seatMap [][]string

func (seats seatMap) print() {
	fmt.Println()
	for _, row := range seats {
		for _, cell := range row {
			fmt.Printf("%v", cell)
		}
		fmt.Println()
	}
}

func (seats seatMap) equals(other seatMap) bool {
	if len(seats) != len(other) {
		fmt.Println("maps are not equal! . len(seats) != len(other)")
		seats.print()
		other.print()
		return false
	}

	for i := range seats {
		row := seats[i]
		otherRow := other[i]

		if len(row) != len(otherRow) {
			fmt.Printf("%v\n not equal \n%v. len(row) != len(otherRow)\n", row, otherRow)
			seats.print()
			other.print()
			return false
		}
		for j := range seats[i] {
			if row[j] != otherRow[j] {
				fmt.Printf("%v\n not equal \n%v\n", row, otherRow)
				seats.print()
				other.print()
				return false
			}
		}
	}
	return true
}

func parseToMap(data []string) seatMap {
	seats := seatMap{}
	rowLen := 0
	for _, line := range data {
		line = strings.TrimSpace(line)
		row := strings.Split(line, "")
		if rowLen == 0 {
			rowLen = len(row)
		}
		if rowLen != len(row) {
			panic("wrong row size!")
		}
		seats = append(seats, row)
	}

	//seats.print()
	return seats
}

func readFile(filename string) seatMap {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := []string{}
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return parseToMap(data)
}

func getNumberOfOccupiedAround(rowNum int, colNum int, seats seatMap) int {
	nubmerOfOccupied := 0
	for r := rowNum - 1; r <= rowNum+1; r++ {
		for c := colNum - 1; c <= colNum+1; c++ {
			if r < 0 || c < 0 || r >= len(seats) || c >= len(seats[r]) {
				continue
			}
			if r == rowNum && c == colNum {
				continue
			}
			if seats[r][c] == "#" {
				nubmerOfOccupied++
			}
		}
	}

	return nubmerOfOccupied
}

func getTotalNumberOfOccupiedSeats(seats seatMap) int {
	nubmerOfOccupied := 0
	for rowNum := range seats {
		for colNum := range seats[rowNum] {
			if seats[rowNum][colNum] == "#" {
				nubmerOfOccupied++
			}
		}
	}

	return nubmerOfOccupied
}

func isGoodToSitTask1(rowNum int, colNum int, seats seatMap) bool {
	if seats[rowNum][colNum] != "L" {
		panic("wrong seat!")
	}

	if getNumberOfOccupiedAround(rowNum, colNum, seats) == 0 {
		return true
	}
	return false
}

func isTooBusyTask1(rowNum int, colNum int, seats seatMap) bool {
	if seats[rowNum][colNum] != "#" {
		panic("wrong seat!")
	}

	if getNumberOfOccupiedAround(rowNum, colNum, seats) >= 4 {
		return true
	}

	return false
}

func doIterationTask1(seats seatMap) seatMap {
	newSeatMap := seatMap{}
	for rowNum := range seats {
		newRow := []string{}
		for colNum := range seats[rowNum] {
			switch seats[rowNum][colNum] {
			case ".":
				newRow = append(newRow, ".")
				continue
			case "L":
				if isGoodToSitTask1(rowNum, colNum, seats) {
					newRow = append(newRow, "#")
				} else {
					newRow = append(newRow, "L")
				}
				continue
			case "#":
				if isTooBusyTask1(rowNum, colNum, seats) {
					newRow = append(newRow, "L")
				} else {
					newRow = append(newRow, "#")
				}
				continue
			default:
				panic("unexpected char")
			}
		}
		newSeatMap = append(newSeatMap, newRow)
	}
	return newSeatMap
}

func task1(seats seatMap) int {
	numberOfOccupiedSeats := 0
	iteration := 0
	for {
		iteration++
		fmt.Printf("Iteration %v\n", iteration)
		seats = doIterationTask1(seats)
		newNumber := getTotalNumberOfOccupiedSeats(seats)
		if newNumber == numberOfOccupiedSeats {
			break
		}
		numberOfOccupiedSeats = newNumber
	}
	return numberOfOccupiedSeats
}

func getNumberOfVisibleOccupied(rowNum int, colNum int, seats seatMap) int {
	nubmerOfOccupied := 0
	offsetLim := int(math.Max(float64(len(seats)), float64(len(seats[0]))))
	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum + offset
		c := colNum
		if r >= len(seats) {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum - offset
		c := colNum
		if r < 0 {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum
		c := colNum + offset
		if c >= len(seats[0]) {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum
		c := colNum - offset
		if c < 0 {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum - offset
		c := colNum - offset
		if c < 0 || r < 0 {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum + offset
		c := colNum - offset
		if c < 0 || r >= len(seats) {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum - offset
		c := colNum + offset
		if r < 0 || c >= len(seats[0]) {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	for offset := 1; offset < offsetLim; offset++ {
		r := rowNum + offset
		c := colNum + offset
		if r >= len(seats) || c >= len(seats[0]) {
			break
		}
		cell := seats[r][c]
		if cell == "." {
			continue
		}
		if cell == "#" {
			nubmerOfOccupied++
			break
		}
		break
	}

	return nubmerOfOccupied
}

func isGoodToSitTask2(rowNum int, colNum int, seats seatMap) bool {
	if seats[rowNum][colNum] != "L" {
		panic("wrong seat!")
	}

	if getNumberOfVisibleOccupied(rowNum, colNum, seats) == 0 {
		return true
	}
	return false
}

func isTooBusyTask2(rowNum int, colNum int, seats seatMap) bool {
	if seats[rowNum][colNum] != "#" {
		panic("wrong seat!")
	}

	if getNumberOfVisibleOccupied(rowNum, colNum, seats) >= 5 {
		return true
	}

	return false
}

func doIterationTask2(seats seatMap) seatMap {
	newSeatMap := seatMap{}
	for rowNum := range seats {
		newRow := []string{}
		for colNum := range seats[rowNum] {
			switch seats[rowNum][colNum] {
			case ".":
				newRow = append(newRow, ".")
				continue
			case "L":
				if isGoodToSitTask2(rowNum, colNum, seats) {
					newRow = append(newRow, "#")
				} else {
					newRow = append(newRow, "L")
				}
				continue
			case "#":
				if isTooBusyTask2(rowNum, colNum, seats) {
					newRow = append(newRow, "L")
				} else {
					newRow = append(newRow, "#")
				}
				continue
			default:
				panic("unexpected char")
			}
		}
		newSeatMap = append(newSeatMap, newRow)
	}
	return newSeatMap
}

func task2(seats seatMap) int {
	numberOfOccupiedSeats := 0
	iteration := 0
	for {
		iteration++
		fmt.Printf("Iteration %v\n", iteration)
		seats = doIterationTask2(seats)
		newNumber := getTotalNumberOfOccupiedSeats(seats)
		if newNumber == numberOfOccupiedSeats {
			break
		}
		numberOfOccupiedSeats = newNumber
	}
	return numberOfOccupiedSeats
}

// task1: 2324
// task2: 0

func main() {
	input := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(input))
	fmt.Printf("Task2: %v\n", task2(input))
}
