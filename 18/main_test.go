package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	expected := 71

	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_2(t *testing.T) {
	input := "1 + (2 * 3) + (4 * (5 + 6))"
	expected := 51
	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_3(t *testing.T) {
	input := "2 * 3 + (4 * 5)"
	expected := 26
	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_4(t *testing.T) {
	input := "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	expected := 437
	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_5(t *testing.T) {
	input := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	expected := 12240

	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_6(t *testing.T) {
	input := "((2 + 1) * (2 + 1) + 2) + 1"
	expected := 12
	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task1_7(t *testing.T) {
	input := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	expected := 13632
	res := evaluateTask1(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_RPL_Task1_1(t *testing.T) {
	input := "1 + 2"
	expected := "12+"

	res, _ := toRPLTask1(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task1_2(t *testing.T) {
	input := "2 * 3 + (4 * 5)"
	expected := "23*45*+"

	res, _ := toRPLTask1(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task1_3(t *testing.T) {
	input := "((2 + 1) * (2 + 1) + 2) + 1"
	expected := "21+21+*2+1+"

	res, _ := toRPLTask1(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_1(t *testing.T) {
	input := "3*1 + 2"
	expected := "312+*"

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_2(t *testing.T) {
	input := "2 * 3 + (4 * 5)"
	expected := "2345*+*"

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_3(t *testing.T) {
	input := "(1*3 + 2)"
	expected := "132+*"
	// ((2*1)+(2*1)*2)*1
	// 21*21*2*+1*
	// 21+21+2+*1+

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_4(t *testing.T) {
	input := "(2) * (2) + 2"
	expected := "222+*"
	// ((2*1)+(2*1)*2)*1
	// 21*21*2*+1*
	// 21+21+2+*1+

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_5(t *testing.T) {
	input := "((2 + 1) * (2 + 1) + 2) + 1"
	expected := "21+21+2+*1+"
	// ((2*1)+(2*1)*2)*1
	// 21*21*2*+1*
	// 21+21+2+*1+

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_RPL_Task2_6(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	expected := "12+34+*56+*"
	// ((2*1)+(2*1)*2)*1
	// 21*21*2*+1*
	// 21+21+2+*1+

	res, _ := toRPLTask2(parseLine(input), 0)
	assert.Equal(t, expected, res)
}

func Test_Task2_1(t *testing.T) {
	input := "1 + 2 * 3 + 4 * 5 + 6"
	expected := 231

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task2_2(t *testing.T) {
	input := "1 + (2 * 3) + (4 * (5 + 6))"
	expected := 51

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task2_3(t *testing.T) {
	input := "2 * 3 + (4 * 5)"
	expected := 46

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task2_4(t *testing.T) {
	input := "5 + (8 * 3 + 9 + 3 * 4 * 3)"
	expected := 1445

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task2_5(t *testing.T) {
	input := "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"
	expected := 669060

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}

func Test_Task2_6(t *testing.T) {
	input := "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"
	expected := 23340

	res := evaluateTask2(parseLine(input))
	assert.Equal(t, expected, res)
}
