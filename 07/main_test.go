package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 4

	// act
	res := task1(input)

	assert.Equal(t, expected, res)
}

func Test_Task2_bags_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	color := "dotted black"
	expected := int64(0)

	// act
	res := bagContainsInside(color, input, map[string]int64{})

	assert.Equal(t, expected, res)
}

func Test_Task2_bags_2(t *testing.T) {
	input := readFile("input_test_task1.txt")
	color := "vibrant plum"
	expected := int64(11)

	// act
	res := bagContainsInside(color, input, map[string]int64{})

	assert.Equal(t, expected, res)
}

func Test_Task2_bags_3(t *testing.T) {
	input := readFile("input_test_task1.txt")
	color := "dark olive"
	expected := int64(7)

	// act
	res := bagContainsInside(color, input, map[string]int64{})

	assert.Equal(t, expected, res)
}

func Test_Task2_bags_4(t *testing.T) {
	input := readFile("input_test_task1.txt")
	color := "shiny gold"
	expected := int64(32)

	// act
	res := bagContainsInside(color, input, map[string]int64{})

	assert.Equal(t, expected, res)
}

func Test_Task2_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := int64(32)

	// act
	res := task2(input)

	assert.Equal(t, expected, res)
}

func Test_Task2_2(t *testing.T) {
	input := readFile("input_test_task2.txt")
	expected := int64(126)

	// act
	res := task2(input)

	assert.Equal(t, expected, res)
}
