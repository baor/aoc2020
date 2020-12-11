package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_0(t *testing.T) {
	sMap := readFile("input_test.txt")
	expected := `L.LL.LL.LL
		LLLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLLL
		L.LLLLLL.L
		L.LLLLL.LL`

	expectedMap := parseToMap(strings.Split(expected, "\n"))
	assert.True(t, sMap.equals(expectedMap))
}

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test.txt")
	expected := 37

	assert.Equal(t, expected, task1(input))
}

func Test_Task1_2(t *testing.T) {
	sMap := readFile("input_test.txt")
	expected := []string{
		`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
		`#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##`,
		`#.##.L#.##
#L###LL.L#
L.#.#..#..
#L##.##.L#
#.##.LL.LL
#.###L#.##
..#.#.....
#L######L#
#.LL###L.L
#.#L###.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.L.L..#..
#LLL.##.L#
#.LL.LL.LL
#.LL#L#.##
..L.L.....
#L#LLLL#L#
#.LLLLLL.L
#.#L#L#.##`,
		`#.#L.L#.##
#LLL#LL.L#
L.#.L..#..
#L##.##.L#
#.#L.LL.LL
#.#L#L#.##
..L.L.....
#L#L##L#L#
#.LLLLLL.L
#.#L#L#.##`}

	for i, val := range expected {
		fmt.Printf("Iteration %v\n", i)
		expectedMap := parseToMap(strings.Split(val, "\n"))
		sMap = doIterationTask1(sMap)
		assert.True(t, sMap.equals(expectedMap))
	}
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test.txt")
	expected := 26

	assert.Equal(t, expected, task2(input))
}

func Test_Task2_Occu1(t *testing.T) {
	input := `.......#.
	...#.....
	.#.......
	.........
	..#L....#
	....#....
	.........
	#........
	...#.....`

	expectedOccupied := 8

	sMap := parseToMap(strings.Split(input, "\n"))
	assert.Equal(t, expectedOccupied, getNumberOfVisibleOccupied(4, 3, sMap))
}

func Test_Task2_Occu2(t *testing.T) {
	input := `.##.##.
	#.#.#.#
	##...##
	...L...
	##...##
	#.#.#.#
	.##.##.`

	expectedOccupied := 0

	sMap := parseToMap(strings.Split(input, "\n"))
	assert.Equal(t, expectedOccupied, getNumberOfVisibleOccupied(3, 3, sMap))
}

func Test_Task2_Occu3(t *testing.T) {
	input := `.............
	.L.L.#.#.#.#.
	.............`

	expectedOccupied := 0

	sMap := parseToMap(strings.Split(input, "\n"))
	assert.Equal(t, expectedOccupied, getNumberOfVisibleOccupied(1, 1, sMap))
}

func Test_Task2_3(t *testing.T) {
	sMap := readFile("input_test.txt")
	expected := []string{
		`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`,
		`#.LL.LL.L#
		#LLLLLL.LL
		L.L.L..L..
		LLLL.LL.LL
		L.LL.LL.LL
		L.LLLLL.LL
		..L.L.....
		LLLLLLLLL#
		#.LLLLLL.L
		#.LLLLL.L#`,
		`#.L#.##.L#
		#L#####.LL
		L.#.#..#..
		##L#.##.##
		#.##.#L.##
		#.#####.#L
		..#.#.....
		LLL####LL#
		#.L#####.L
		#.L####.L#`,
		`#.L#.L#.L#
		#LLLLLL.LL
		L.L.L..#..
		##LL.LL.L#
		L.LL.LL.L#
		#.LLLLL.LL
		..L.L.....
		LLLLLLLLL#
		#.LLLLL#.L
		#.L#LL#.L#`,
		`#.L#.L#.L#
		#LLLLLL.LL
		L.L.L..#..
		##L#.#L.L#
		L.L#.#L.L#
		#.L####.LL
		..#.#.....
		LLL###LLL#
		#.LLLLL#.L
		#.L#LL#.L#`,
		`#.L#.L#.L#
	#LLLLLL.LL
	L.L.L..#..
	##L#.#L.L#
	L.L#.LL.L#
	#.LLLL#.LL
	..#.L.....
	LLL###LLL#
	#.LLLLL#.L
	#.L#LL#.L#`}

	for i, val := range expected {
		fmt.Printf("Iteration %v\n", i)
		expectedMap := parseToMap(strings.Split(val, "\n"))
		sMap = doIterationTask2(sMap)
		assert.True(t, sMap.equals(expectedMap))
	}
}
