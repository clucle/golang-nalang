package snakegame

import (
	"os"
	"os/exec"
	"runtime"
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
