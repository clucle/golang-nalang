package main

import (
	"fmt"
	"time"

	"github.com/clucle/golang-nalang/src/console/snake-game/snakegame"
	"github.com/eiannone/keyboard"
)

func main() {
	snakegame.ClearScreen()

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

	go func() {
		elpasedTime := 0
		for {
			select {
			case <-done:
				return

			case <-ticker.C:
				elpasedTime += 1
				snakegame.ClearScreen()
				// snakegame.Display()
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

	snakegame.ClearScreen()
}
