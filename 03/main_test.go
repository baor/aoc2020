package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test.txt")
	expectedResponse := 7

	// act
	res := task1(input)

	assert.Equal(t, expectedResponse, res)
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test.txt")
	expectedResponse := 336

	// act
	res := task2(input)

	assert.Equal(t, expectedResponse, res)
}
