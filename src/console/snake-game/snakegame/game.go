package snakegame

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

const boardSize = 13

type KeyState int

const (
	KeyUp KeyState = 1 + iota
	KeyDown
	KeyLeft
	KeyRight
)

type Game struct {
	ticker      *time.Ticker
	done        chan bool
	board       [boardSize][boardSize]int
	keyState    KeyState
	elapsedTime int
}

func (game *Game) Init() {
	game.ticker = time.NewTicker(time.Millisecond * 160)
	game.done = make(chan bool)
	game.elapsedTime = 0

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			game.board[row][col] = 0
		}
	}
}

func (game *Game) Run() {
	// set key events
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	go func() {
		for {
			select {
			case <-game.done:
				return

			case <-game.ticker.C:
				game.elapsedTime++
				ClearScreen()
				game.Display()
				fmt.Printf("%d", game.keyState)
			}
		}
	}()

	game.keyState = KeyUp

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			game.ticker.Stop()
			game.done <- true
			break
		} else if event.Key == keyboard.KeyArrowDown || event.Key == 65516 {
			game.keyState = KeyDown
		} else if event.Key == keyboard.KeyArrowUp || event.Key == 65517 {
			game.keyState = KeyUp
		} else if event.Key == keyboard.KeyArrowLeft || event.Key == 65515 {
			game.keyState = KeyLeft
		} else if event.Key == keyboard.KeyArrowRight || event.Key == 65514 {
			game.keyState = KeyRight
		}
	}

}

func (game *Game) Display() {
	display := fmt.Sprintf("Elapsed Time : %d\n", game.elapsedTime)
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			display += "▢ "
		}
		display += "\n"
	}
	fmt.Println(display)
}
