package main

import "fmt"

func main() {
	game := NewGame()
	game.Add(5)
	game.Add(7)
	fmt.Println("Game score: ", game.Score())
}
