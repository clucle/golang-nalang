package snakegame

import "fmt"

const boardSize = 13

type Game struct {
	board [boardSize][boardSize]int
}

func Display(game Game) {
	var ret string
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			ret += "▢ "
		}
		ret += "\n"
	}
	fmt.Println(ret)
}
