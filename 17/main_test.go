package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 112

	assert.Equal(t, expected, task1(input))
}

func Test_Task1_3(t *testing.T) {
	cube := readFile("input_test_task1.txt")
	zMin := -1
	expected := []int{3, 5, 3}

	cube = doIteration(cube, 0)
	cube.printAll()
	for i, val := range expected {
		fmt.Printf("Iteration %v\n", i)
		assert.Equal(t, val, cube.numberOfActive(0, zMin+i))
	}
}

func Test_Task1_4(t *testing.T) {
	cube := readFile("input_test_task1.txt")
	zMin := -2
	expected := []int{1, 5, 9, 5, 1}

	cube = doIteration(cube, 0)
	cube = doIteration(cube, 0)
	cube.printAll()
	for i, val := range expected {
		fmt.Printf("Iteration %v\n", i)
		assert.Equal(t, val, cube.numberOfActive(0, zMin+i))
	}
}

func Test_Task1_5(t *testing.T) {
	cube := readFile("input_test_task1.txt")
	zMin := -2
	expected := []int{5, 10, 8, 10, 5}

	cube = doIteration(cube, 0)
	cube = doIteration(cube, 0)
	cube = doIteration(cube, 0)
	cube.printAll()
	for i, val := range expected {
		fmt.Printf("Iteration %v\n", i)
		assert.Equal(t, val, cube.numberOfActive(0, zMin+i))
	}
}

func Test_Task2_1(t *testing.T) {
	cube := readFile("input_test_task1.txt")
	cube = doIteration(cube, 1)
	cube.printAll()
	assert.Equal(t, 3, cube.numberOfActive(-1, -1))
	assert.Equal(t, 3, cube.numberOfActive(-1, 0))
	assert.Equal(t, 3, cube.numberOfActive(-1, 1))

	assert.Equal(t, 3, cube.numberOfActive(0, -1))
	assert.Equal(t, 5, cube.numberOfActive(0, 0))
	assert.Equal(t, 3, cube.numberOfActive(0, 1))

	assert.Equal(t, 3, cube.numberOfActive(1, -1))
	assert.Equal(t, 3, cube.numberOfActive(1, 0))
	assert.Equal(t, 3, cube.numberOfActive(1, 1))
}

func Test_Task2_2(t *testing.T) {
	cube := readFile("input_test_task1.txt")
	cube = doIteration(cube, 1)
	cube = doIteration(cube, 1)
	cube.printAll()
	assert.Equal(t, 1, cube.numberOfActive(-2, -2))
	assert.Equal(t, 0, cube.numberOfActive(-2, -1))
	assert.Equal(t, 14, cube.numberOfActive(-2, 0))
	assert.Equal(t, 0, cube.numberOfActive(-2, 1))
	assert.Equal(t, 1, cube.numberOfActive(-2, 2))

	assert.Equal(t, 0, cube.numberOfActive(-1, -2))
	assert.Equal(t, 0, cube.numberOfActive(-1, -1))
	assert.Equal(t, 0, cube.numberOfActive(-1, 0))
	assert.Equal(t, 0, cube.numberOfActive(-1, 1))
	assert.Equal(t, 0, cube.numberOfActive(-1, 2))

	assert.Equal(t, 14, cube.numberOfActive(0, -2))
	assert.Equal(t, 0, cube.numberOfActive(0, -1))
	assert.Equal(t, 0, cube.numberOfActive(0, 0))
	assert.Equal(t, 0, cube.numberOfActive(0, 1))
	assert.Equal(t, 14, cube.numberOfActive(0, 2))

	assert.Equal(t, 0, cube.numberOfActive(1, -2))
	assert.Equal(t, 0, cube.numberOfActive(1, -1))
	assert.Equal(t, 0, cube.numberOfActive(1, 0))
	assert.Equal(t, 0, cube.numberOfActive(1, 1))
	assert.Equal(t, 0, cube.numberOfActive(1, 2))

	assert.Equal(t, 1, cube.numberOfActive(2, -2))
	assert.Equal(t, 0, cube.numberOfActive(2, -1))
	assert.Equal(t, 14, cube.numberOfActive(2, 0))
	assert.Equal(t, 0, cube.numberOfActive(2, 1))
	assert.Equal(t, 1, cube.numberOfActive(2, 2))
}

func Test_Task2_L(t *testing.T) {
	input := readFile("input_test_task1.txt")
	expected := 848

	assert.Equal(t, expected, task2(input))
}
