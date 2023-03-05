package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/eiannone/keyboard"
)

func ClearScreen() {
	osString := runtime.GOOS
	if osString == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osString == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

const boardSize = 13

func GetBoardStr(board [boardSize][boardSize]int) string {
	var ret string
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			ret += "▢ "
		}
		ret += "\n"
	}
	return ret
}

func main() {
	ticker := time.NewTicker(time.Millisecond * 200)
	done := make(chan bool)

	// set key events
	keysEvents, err := keyboard.GetKeys(10)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	type KeyState int
	const (
		KeyUp KeyState = 1 + iota
		KeyDown
		KeyLeft
		KeyRight
	)
	var keyState = KeyUp

	// set board
	var board [boardSize][boardSize]int

	go func() {
		elpasedTime := 0
		for {
			select {
			case <-done:
				return

			case <-ticker.C:
				elpasedTime += 1
				ClearScreen()
				display := fmt.Sprintf("Elapsed Time : %d\n", elpasedTime)
				display += GetBoardStr(board)
				fmt.Println(display)
				fmt.Printf("%d", keyState)
			}
		}
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			ticker.Stop()
			done <- true
			break
		} else if event.Key == keyboard.KeyArrowDown || event.Key == 65516 {
			keyState = KeyDown
		} else if event.Key == keyboard.KeyArrowUp || event.Key == 65517 {
			keyState = KeyUp
		} else if event.Key == keyboard.KeyArrowLeft || event.Key == 65515 {
			keyState = KeyLeft
		} else if event.Key == keyboard.KeyArrowRight || event.Key == 65514 {
			keyState = KeyRight
		}
	}

	ClearScreen()
}
