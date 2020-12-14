package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	time, busses := readFile("input_test.txt")
	expected := 295

	assert.Equal(t, expected, task1(time, busses))
}

func Test_Task2_1(t *testing.T) {
	_, busses := readFile("input_test.txt")
	expected := 1068781

	assert.Equal(t, expected, task2(busses, 1))
}

func Test_Task2_2(t *testing.T) {
	busses := parseSchedule("17,x,13,19")
	expected := 3417

	assert.Equal(t, expected, task2(busses, 1))
}

func Test_Task2_3(t *testing.T) {
	busses := parseSchedule("67,7,59,61")
	expected := 754018

	assert.Equal(t, expected, task2(busses, 1))
}

func Test_Task2_4(t *testing.T) {
	busses := parseSchedule("67,x,7,59,61")
	expected := 779210

	assert.Equal(t, expected, task2(busses, 1))
}

func Test_Task2_5(t *testing.T) {
	busses := parseSchedule("67,7,x,59,61")
	expected := 1261476

	assert.Equal(t, expected, task2(busses, 1))
}

func Test_Task2_6(t *testing.T) {
	busses := parseSchedule("1789,37,47,1889")
	expected := 1202161486

	assert.Equal(t, expected, task2(busses, 1))
}
