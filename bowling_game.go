package main

const Strike = 10
const NumberOfFrames = 10

type Frame struct {
	FirstThrow, SecondThrow int
}

func NewFrame(firstThrow, secondThrow int) Frame {
	return Frame{FirstThrow: firstThrow, SecondThrow: secondThrow}
}

func (frame Frame) Score(next, followingNext Frame) int {
	if frame.IsStrike() {
		if next.IsStrike() {
			return frame.FirstThrow + next.FirstThrow + followingNext.FirstThrow
		} else {
			return frame.FirstThrow + next.Sum()
		}
	} else if frame.IsSpare() {
		return frame.FirstThrow + frame.SecondThrow + next.FirstThrow
	} else {
		return frame.FirstThrow + frame.SecondThrow
	}
}

func (frame Frame) IsStrike() bool {
	return frame.FirstThrow == Strike
}

func (frame Frame) IsSpare() bool {
	return frame.Sum() == Strike && !frame.IsStrike()
}

func (frame Frame) Sum() int {
	return frame.FirstThrow + frame.SecondThrow
}

type Game struct {
	currentFrame      int
	frames            [12]Frame
	currentFrameEnded bool
	currentFrameValue int
}

func NewGame() *Game {
	return &Game{currentFrame: 0, currentFrameEnded: true, currentFrameValue: 0}
}

func (game *Game) Add(pins int) {
	if pins < 10 {
		if game.currentFrameEnded {
			game.currentFrameEnded = false
			game.currentFrameValue = pins
		} else {
			game.frames[game.currentFrame] = NewFrame(game.currentFrameValue, pins)
			game.currentFrameEnded = true
		}
	}
}

func (game Game) Score() int {
	result := 0

	for index := 0; index < NumberOfFrames; index++ {
		frame := game.frames[index]
		next := game.frames[index+1]
		followingNext := game.frames[index+2]
		result += frame.Score(next, followingNext)
	}

	return result
}
