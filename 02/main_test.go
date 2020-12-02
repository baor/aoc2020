package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test.txt")
	expectedResponse := 2

	// act
	res := task1(input)

	assert.Equal(t, expectedResponse, res)
}

func Test_Task2_1(t *testing.T) {
	assert.True(t, task2IsValid(lineWithPassword{
		minNum:   1,
		maxNum:   3,
		char:     "a",
		password: "abcde",
	}))
	assert.False(t, task2IsValid(lineWithPassword{
		minNum:   1,
		maxNum:   3,
		char:     "b",
		password: "cdefg",
	}))
	assert.False(t, task2IsValid(lineWithPassword{
		minNum:   2,
		maxNum:   9,
		char:     "c",
		password: "ccccccccc",
	}))
}

func Test_Task2_2(t *testing.T) {
	input := readFile("input_test.txt")
	expectedResponse := 1

	// act
	res := task2(input)

	assert.Equal(t, expectedResponse, res)
}
