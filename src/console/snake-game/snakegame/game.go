package snakegame

import (
	"fmt"
	"math/rand"
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

type Position struct {
	row, col int
}

func IsValid(position Position) bool {
	if position.row < 0 || position.col < 0 || position.row >= boardSize || position.col >= boardSize {
		return false
	}
	return true
}

type Game struct {
	ticker   *time.Ticker
	done     chan bool
	exit     chan bool
	keyState KeyState

	board [boardSize][boardSize]int
	snake Snake
	apple Apple

	elapsedTime int
	score       int
}

func (game *Game) Init() {
	game.ticker = time.NewTicker(time.Millisecond * 160)
	game.done = make(chan bool)
	game.exit = make(chan bool)
	game.elapsedTime = 0
	game.score = 0

	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			game.board[row][col] = 0
		}
	}

	game.snake.Init(boardSize)
	game.apple = Apple{Position{2, boardSize / 2}}

	game.board[game.snake.body[0].row][game.snake.body[0].col] = 1
	game.board[game.apple.body.row][game.apple.body.col] = 2
}

func (game *Game) Run() {
	// set key events
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case <-game.exit:
				game.done <- true
				return
			case <-game.ticker.C:
				game.elapsedTime++
				ClearScreen()
				if !game.Update() {
					game.Display()
					game.done <- true
					return
				}
				game.Display()
			default:
			}
		}
	}()

	game.keyState = KeyUp

	for {
		select {
		case <-game.done:
			return
		default:
		}

		select {
		case event := <-keysEvents:
			if event.Err != nil {
				panic(event.Err)
			}
			if event.Key == keyboard.KeyEsc {
				game.ticker.Stop()
				game.exit <- true
			} else if event.Key == keyboard.KeyArrowDown || event.Key == 65516 {
				game.keyState = KeyDown
			} else if event.Key == keyboard.KeyArrowUp || event.Key == 65517 {
				game.keyState = KeyUp
			} else if event.Key == keyboard.KeyArrowLeft || event.Key == 65515 {
				game.keyState = KeyLeft
			} else if event.Key == keyboard.KeyArrowRight || event.Key == 65514 {
				game.keyState = KeyRight
			}
		default:
		}
	}
}

func (game *Game) Update() bool {
	row, col := game.snake.body[0].row, game.snake.body[0].col

	switch keyState := game.keyState; keyState {
	case KeyDown:
		row++
	case KeyUp:
		row--
	case KeyLeft:
		col--
	case KeyRight:
		col++
	}

	if !IsValid(Position{row, col}) || game.board[row][col] == 1 {
		return false
	}

	body := []Position{{row, col}}

	game.board[row][col] = 1

	if row == game.apple.body.row && col == game.apple.body.col {
		game.score++
		game.snake.body = append(body, game.snake.body[:]...)

		position, valid := game.GenerateApple()
		if !valid {
			return false
		}

		game.apple.body = position
		game.board[position.row][position.col] = 2

	} else {
		tailIndex := len(game.snake.body) - 1
		tailRow, tailCol := game.snake.body[tailIndex].row, game.snake.body[tailIndex].col
		game.board[tailRow][tailCol] = 0

		game.snake.body = append(body, game.snake.body[:tailIndex]...)
	}

	return true
}

func (game *Game) GenerateApple() (Position, bool) {
	candidate := []Position{}
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if game.board[row][col] == 0 {
				candidate = append(candidate, Position{row, col})
			}
		}
	}

	if len(candidate) == 0 {
		return Position{0, 0}, false
	}

	src := rand.NewSource(time.Now().UnixNano())
	random := rand.New(src)
	index := random.Intn(len(candidate))

	return candidate[index], true
}

func (game *Game) Display() {
	var Reset = "\033[0m"
	var Red = "\033[31m"
	var White = "\033[97m"
	var Green = "\033[32m"

	display := White
	display += fmt.Sprintf("Elapsed Time : %d\n", game.elapsedTime)
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if game.board[row][col] == 0 {
				display += "▢ "
			} else if game.board[row][col] == 1 {
				display += Reset
				display += Green
				display += "◯ "
				display += Reset
			} else if game.board[row][col] == 2 {
				display += Reset
				display += Red
				display += "▣ "
				display += Reset
			}
		}
		display += "\n"
	}
	display += fmt.Sprintf("Score : %d\n", game.score)
	display += Reset
	fmt.Println(display)
}
