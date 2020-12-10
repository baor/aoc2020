package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	preamble := 5
	expected := 127

	// act
	res := getWrongNumber(input, preamble)

	assert.Equal(t, expected, res)
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	preamble := 5
	expected := 62

	// act
	assert.Equal(t, expected, getXMASNumber(getWrongNumber(input, preamble), input))
}
