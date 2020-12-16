package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Task1_1(t *testing.T) {
	target := 2020
	assert.Equal(t, 436, task1("0,3,6", target))
	assert.Equal(t, 1, task1("1,3,2", target))
	assert.Equal(t, 10, task1("2,1,3", target))
	assert.Equal(t, 27, task1("1,2,3", target))
	assert.Equal(t, 78, task1("2,3,1", target))
	assert.Equal(t, 438, task1("3,2,1", target))
	assert.Equal(t, 1836, task1("3,1,2", target))
}

// func Test_Task2_1(t *testing.T) {
// 	target := 30000000
// 	assert.Equal(t, 175594, task1("0,3,6", target))
// 	assert.Equal(t, 2578, task1("1,3,2", target))
// 	assert.Equal(t, 3544142, task1("2,1,3", target))
// 	assert.Equal(t, 261214, task1("1,2,3", target))
// 	assert.Equal(t, 6895259, task1("2,3,1", target))
// 	assert.Equal(t, 18, task1("3,2,1", target))
// 	assert.Equal(t, 362, task1("3,1,2", target))
// }
