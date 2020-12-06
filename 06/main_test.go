package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 11

	// act
	res := task1(input)

	assert.Equal(t, expected, res)
}

func Test_Task1_CreateAnswer_1(t *testing.T) {
	input := `ab
	ac`
	expected := answers{"a": 2, "b": 1, "c": 1}

	// act
	a := createAnswersOfTheGroup(input)

	assert.Equal(t, expected, a)
}

func Test_Task3_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 6

	// act
	res := task2(input)

	assert.Equal(t, expected, res)
}
