package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test.txt")
	expected := 25

	assert.Equal(t, expected, task1(input))
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test.txt")
	expected := 286

	assert.Equal(t, expected, task2(input))
}
