package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	input := "FBFBBFFRLR"
	expected := seat{
		row:    44,
		column: 5,
		seatID: 357,
	}

	// act
	s := createSeat(input)

	assert.Equal(t, expected, s)
}

func Test_Task1_2(t *testing.T) {
	input := "BFFFBBFRRR"
	expected := seat{
		row:    70,
		column: 7,
		seatID: 567,
	}

	// act
	s := createSeat(input)

	assert.Equal(t, expected, s)
}

func Test_Task1_3(t *testing.T) {
	input := "FFFBBBFRRR"
	expected := seat{
		row:    14,
		column: 7,
		seatID: 119,
	}

	// act
	s := createSeat(input)

	assert.Equal(t, expected, s)
}

func Test_Task1_4(t *testing.T) {
	input := "BBFFBBFRLL"
	expected := seat{
		row:    102,
		column: 4,
		seatID: 820,
	}

	// act
	s := createSeat(input)

	assert.Equal(t, expected, s)
}
