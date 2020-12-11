package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_1.txt")
	expected := 22

	assert.Equal(t, expected, input[len(input)-1])
}

func Test_Task1_2(t *testing.T) {
	input := readFile("input_test_1.txt")
	expected := 35

	assert.Equal(t, expected, task1(input))
}

func Test_Task1_3(t *testing.T) {
	input := readFile("input_test_2.txt")
	expected := 220

	assert.Equal(t, expected, task1(input))
}

func Test_Task1_4(t *testing.T) {
	input := readFile("input.txt")
	expected := 3000

	assert.Equal(t, expected, task1(input))
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test_1.txt")
	expected := 8

	assert.Equal(t, expected, task2(input))
}

func Test_Task2_2(t *testing.T) {
	input := readFile("input_test_2.txt")
	expected := 19208

	assert.Equal(t, expected, task2(input))
}
