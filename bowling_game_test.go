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
		{description: "spare", throws: []int{8, 2, 5, 0}, expectedResult: 20},
		{description: "another spare", throws: []int{7, 3, 3, 0}, expectedResult: 16},
		{description: "spare after spare", throws: []int{7, 3, 7, 3, 0, 0}, expectedResult: 27},
		{description: "spare after spare after spare", throws: []int{7, 3, 7, 3, 7, 3, 0, 0}, expectedResult: 44},
		{description: "stimple strike 1", throws: []int{10, 0, 7}, expectedResult: 24},
		{description: "stimple strike 2", throws: []int{10, 2, 7}, expectedResult: 28},
		{description: "strike after spare", throws: []int{2, 8, 10, 3, 4}, expectedResult: 44},
		{description: "strike after strike", throws: []int{10, 10, 3, 4}, expectedResult: 47},
		{description: "full game example", throws: []int{1, 4, 4, 5, 6, 4, 5, 5, 10, 0, 1, 7, 3, 6, 4, 10, 2, 7}, expectedResult: 125},
	}

	for _, testCase := range testCases {

		game := NewGame()

		for _, frame := range testCase.throws {
			game.Add(frame)
		}
		assert.Equal(t, testCase.expectedResult, game.Score(), testCase.description)
	}
}

func TestFrameIsStrike(t *testing.T) {
	frame := Frame{FirstThrow: 10, SecondThrow: 0}
	assert.True(t, frame.IsStrike())
}

func TestFrameIsSpare(t *testing.T) {
	testCases := []Frame{
		{FirstThrow: 1, SecondThrow: 9},
		{FirstThrow: 3, SecondThrow: 7},
		{FirstThrow: 8, SecondThrow: 2},
	}

	for _, frame := range testCases {
		assert.True(t, frame.IsSpare())
	}
}

func TestFrameIsNormalCase(t *testing.T) {
	testCases := []Frame{
		{FirstThrow: 1, SecondThrow: 3},
		{FirstThrow: 2, SecondThrow: 4},
		{FirstThrow: 4, SecondThrow: 4},
	}

	for _, frame := range testCases {
		assert.False(t, frame.IsSpare())
		assert.False(t, frame.IsStrike())
	}
}

func TestFrameCalculateScore(t *testing.T) {
	testCases := []testGame{
		{description: "single frame 1", throws: []int{1, 3, 0, 0, 0, 0}, expectedResult: 4},
		{description: "single frame 2", throws: []int{2, 2, 0, 0, 0, 0}, expectedResult: 4},
		{description: "single frame 3", throws: []int{3, 5, 0, 0, 0, 0}, expectedResult: 8},
		{description: "single frame 4", throws: []int{1, 8, 0, 0, 0, 0}, expectedResult: 9},
		{description: "spare 1", throws: []int{1, 9, 0, 0, 0, 0}, expectedResult: 10},
		{description: "spare 2", throws: []int{1, 9, 4, 0, 0, 0}, expectedResult: 14},
		{description: "spare 3", throws: []int{1, 9, 2, 0, 0, 0}, expectedResult: 12},
		{description: "spare 4", throws: []int{1, 9, 10, 0, 0, 0}, expectedResult: 20},
		{description: "strike 1", throws: []int{10, 0, 1, 1, 0, 0}, expectedResult: 12},
		{description: "strike 2", throws: []int{10, 0, 2, 3, 0, 0}, expectedResult: 15},
		{description: "strike 3", throws: []int{10, 0, 10, 0, 10, 0}, expectedResult: 30},
		{description: "strike 4", throws: []int{10, 0, 0, 0, 0, 0}, expectedResult: 10},
		{description: "strike 5", throws: []int{10, 0, 5, 5, 0, 0}, expectedResult: 20},
		{description: "strike 6", throws: []int{10, 0, 7, 2, 0, 0}, expectedResult: 19},
	}

	for _, testCase := range testCases {
		frame := Frame{FirstThrow: testCase.throws[0], SecondThrow: testCase.throws[1]}
		next := Frame{FirstThrow: testCase.throws[2], SecondThrow: testCase.throws[3]}
		followingNext := Frame{FirstThrow: testCase.throws[4], SecondThrow: testCase.throws[5]}

		result := frame.Score(next, followingNext)
		assert.Equal(t, testCase.expectedResult, result, testCase.description)
	}
}
