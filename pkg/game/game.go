package game

import (
	"fmt"
	"log"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

// Run starts the game
func Run() error {
	fmt.Println("In game.Run()")

	runtime.LockOSThread()

	window, err := sdl.CreateWindow("cardshop", 100, 100, 800, 600, sdl.WINDOW_RESIZABLE)
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}

	_ = renderer

	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			processEvent(e)
		}

		if err := renderer.Clear(); err != nil {
			panic(err)
		}
		renderer.SetDrawColor(127, 127, 255, 0)
		renderer.Present()
	}

	sdl.Quit()

	return nil
}

var quit bool

func processEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyboardEvent:
		log.Println(e)
		if e.Keysym.Sym == sdl.K_ESCAPE {
			quit = true
		}
	case *sdl.QuitEvent:
		quit = true
	default:
		//
	}
}
