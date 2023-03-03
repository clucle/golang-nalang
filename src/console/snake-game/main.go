package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {
		for t := range ticker.C {
			fmt.Print("\033[H\033[2J") // clear screen
			fmt.Println("current time", t)
		}
	}()

	input := 0
	fmt.Scanln(&input)
	ticker.Stop()
}
