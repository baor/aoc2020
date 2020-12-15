package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	prog := readFileTask1("input_test_task1.txt")
	expected := uint(165)

	assert.Equal(t, expected, task1(prog))
}

func Test_Task1_2(t *testing.T) {
	f0, f1 := parseMaskTask1("1101")

	assert.Equal(t, "10", fmt.Sprintf("%b", f0))
	assert.Equal(t, "1101", fmt.Sprintf("%b", f1))
}

func Test_Task2_1(t *testing.T) {
	fPairs := parseMaskTask2("1X")

	assert.Equal(t, "1", fmt.Sprintf("%b", fPairs[0].zero))
	assert.Equal(t, "10", fmt.Sprintf("%b", fPairs[0].one))
	assert.Equal(t, "0", fmt.Sprintf("%b", fPairs[1].zero))
	assert.Equal(t, "11", fmt.Sprintf("%b", fPairs[1].one))
}

func Test_Task2_2(t *testing.T) {
	fPairs := parseMaskTask2("0X1001X")

	assert.Equal(t, "100001", fmt.Sprintf("%b", fPairs[0].zero))
	assert.Equal(t, "10010", fmt.Sprintf("%b", fPairs[0].one))
	assert.Equal(t, "1", fmt.Sprintf("%b", fPairs[1].zero))
	assert.Equal(t, "110010", fmt.Sprintf("%b", fPairs[1].one))
	assert.Equal(t, "100000", fmt.Sprintf("%b", fPairs[2].zero))
	assert.Equal(t, "10011", fmt.Sprintf("%b", fPairs[2].one))
	assert.Equal(t, "0", fmt.Sprintf("%b", fPairs[3].zero))
	assert.Equal(t, "110011", fmt.Sprintf("%b", fPairs[3].one))
}

func Test_Task2_3(t *testing.T) {
	fPairs := parseMaskTask2("XX")

	assert.Equal(t, "11", fmt.Sprintf("%b", fPairs[0].zero))
	assert.Equal(t, "0", fmt.Sprintf("%b", fPairs[0].one))
	assert.Equal(t, "1", fmt.Sprintf("%b", fPairs[1].zero))
	assert.Equal(t, "10", fmt.Sprintf("%b", fPairs[1].one))
	assert.Equal(t, "10", fmt.Sprintf("%b", fPairs[2].zero))
	assert.Equal(t, "1", fmt.Sprintf("%b", fPairs[2].one))
	assert.Equal(t, "0", fmt.Sprintf("%b", fPairs[3].zero))
	assert.Equal(t, "11", fmt.Sprintf("%b", fPairs[3].one))
}

func Test_Task2_L(t *testing.T) {
	prog := readFileTask2("input_test_task2.txt")
	expected := uint(208)

	assert.Equal(t, expected, task2(prog))
}
