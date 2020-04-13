package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testGame struct {
	throws         []int
	expectedResult int
	description    string
}

func TestGameCalculateScore(t *testing.T) {
	testCases := []testGame{
		{description: "single frame", throws: []int{1, 4}, expectedResult: 5},
		{description: "another single frame", throws: []int{2, 6}, expectedResult: 8},
	}

	game := NewGame()

	for _, testCase := range testCases {

		for _, frame := range testCase.throws {
			game.Add(frame)
		}
		assert.Equal(t, testCase.expectedResult, game.Score())
	}
}
