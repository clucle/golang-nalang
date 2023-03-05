package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func clearScreen() {
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

func main() {
	ticker := time.NewTicker(time.Millisecond * 100)
	done := make(chan bool)

	go func() {
		elpasedTime := 0
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				elpasedTime += 1
				fmt.Print("\033[H\033[2J")
				fmt.Println("Elapsed Time : ", elpasedTime)
			}
		}
	}()

	input := 0
	fmt.Scanln(&input)
	ticker.Stop()
	done <- true

	clearScreen()
}
