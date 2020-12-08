package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type seat struct {
	row    int
	column int
	seatID int
}

func createSeat(input string) seat {
	rowNumbersBegin := 0
	rowNubmersEnd := 127

	columnNumbersBegin := 0
	columnNumbersEnd := 7
	for i := 0; i < len(input); i++ {
		numberOfRows := rowNubmersEnd - rowNumbersBegin + 1
		shift := numberOfRows / 2
		switch input[i] {
		case 'F':
			rowNubmersEnd -= shift
			continue
		case 'B':
			rowNumbersBegin += shift
			continue
		default:
			if numberOfRows != 1 {
				panic("Expected F and B, got: \"" + string(input[i]) + "\"")
			}
		}
		numberOfcolumns := columnNumbersEnd - columnNumbersBegin + 1
		shift = numberOfcolumns / 2
		switch input[i] {
		case 'L':
			columnNumbersEnd -= shift
			continue
		case 'R':
			columnNumbersBegin += shift
			continue
		default:
			panic("Expected L and R, got: \"" + string(input[i]) + "\"")
		}
	}

	if rowNumbersBegin != rowNubmersEnd {
		panic("wrong rows calculation")
	}
	if columnNumbersBegin != columnNumbersEnd {
		panic("wrong columns calculation")
	}

	return seat{
		row:    rowNumbersBegin,
		column: columnNumbersBegin,
		seatID: rowNumbersBegin*8 + columnNumbersBegin}
}

func readFile(filename string) []seat {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats := []seat{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		seat := createSeat(scanner.Text())
		seats = append(seats, seat)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%+v\n", passports)
	return seats
}

func task1(seats []seat) int {
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].seatID < seats[j].seatID
	})

	//fmt.Printf("%+v\n", seats)

	return seats[len(seats)-1].seatID
}

func createSeatMap() [][]bool {
	seatMap := make([][]bool, 128)
	for i := range seatMap {
		seatMap[i] = make([]bool, 8)
	}

	return seatMap
}

func findMissingSeats(seats []seat) {
	seatMap := createSeatMap()
	for _, seat := range seats {
		seatMap[seat.row][seat.column] = true
	}

	for i, row := range seatMap {
		for j, taken := range row {
			if !taken {
				fmt.Printf("%v:%v, id: %v\n", i, j, i*8+j)
			}
		}
	}
}

func task2(seats []seat) int {
	findMissingSeats(seats)
	return 0
}

// task1: 858
// task2: 557

func main() {
	seats := readFile("input.txt")
	fmt.Printf("Task1: %v\n", task1(seats))
	fmt.Printf("Task2: %v\n", task2(seats))
}
