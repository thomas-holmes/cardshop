package game

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

// Run starts the game
func Run() error {
	fmt.Println("In game.Run()")

	window, err := sdl.CreateWindow("cardshop", 0, 0, 800, 600, 0)
	if err != nil {
		return err
	}

	time.Sleep(2 * time.Second)

	_ = window

	return nil
}
