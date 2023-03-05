package main

import (
	"github.com/clucle/golang-nalang/src/console/snake-game/snakegame"
)

func main() {
	game := &snakegame.Game{}
	game.Init()
	game.Run()
	// snakegame.ClearScreen()
}
