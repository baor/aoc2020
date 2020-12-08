package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 5

	// act
	res := task1(input)

	assert.Equal(t, expected, res)
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 8

	// act
	res := task2(input)

	assert.Equal(t, expected, res)
}
