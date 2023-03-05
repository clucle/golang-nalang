package snakegame

type Snake struct {
	body []Position
}

func (snake *Snake) Init(boardSize int) {
	snake.body = append(snake.body, Position{boardSize / 2, boardSize / 2})
}
