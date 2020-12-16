package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	prog := readFile("input_test_task1.txt")
	expected := 71

	assert.Equal(t, expected, task1(prog))
}

func Test_Task2_1(t *testing.T) {
	d := readFile("input_test_task2.txt")
	expected := map[string]int{"class": 1, "row": 0, "seat": 2}

	assert.Equal(t, expected, identifyMapping(d))
}
