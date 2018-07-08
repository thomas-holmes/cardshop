package game

import (
	"fmt"
	"log"
	"runtime"
	"sort"

	"github.com/veandco/go-sdl2/sdl"
)

// Run starts the game
func Run() error {
	fmt.Println("In game.Run()")

	runtime.LockOSThread()

	window, err := sdl.CreateWindow("cardshop", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 1200, 900, sdl.WINDOW_RESIZABLE)
	if err != nil {
		return err
	}

	renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		return err
	}

	if err := renderer.SetLogicalSize(1200, 900); err != nil {
		return err
	}

	_ = renderer

	cards = append(cards, card{rect: sdl.Rect{X: 100, Y: 200, W: 50, H: 50}, color: sdl.Color{R: 0, G: 0, B: 255, A: 255}})
	cards = append(cards, card{rect: sdl.Rect{X: 400, Y: 200, W: 50, H: 50}, color: sdl.Color{R: 255, G: 0, B: 0, A: 255}})
	cards = append(cards, card{rect: sdl.Rect{X: 300, Y: 350, W: 50, H: 50}, color: sdl.Color{R: 0, G: 255, B: 0, A: 255}})

	for !quit {
		for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
			processEvent(e)
		}

		if err := renderer.Clear(); err != nil {
			panic(err)
		}
		if err := drawBoard(renderer); err != nil {
			return err
		}
		if err := drawCards(renderer); err != nil {
			return err
		}

		renderer.SetDrawColor(127, 127, 255, 255)
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
	case *sdl.MouseButtonEvent:
		if e.Button != sdl.BUTTON_LEFT {
			break
		}
		if e.State == sdl.PRESSED {
			card, ok := checkCards(e.X, e.Y)
			if ok {
				log.Println("Clicked card", card)
				grabCard(card)
			}
		}
		if e.State == sdl.RELEASED {
			releaseCard()
		}
	case *sdl.MouseMotionEvent:
		if err := dragCard(e.XRel, e.YRel); err != nil {
			panic(err)
		}
	case *sdl.QuitEvent:
		quit = true
	default:
		//
	}
}

func drawBoard(renderer *sdl.Renderer) error {
	w, h := renderer.GetLogicalSize()

	r := sdl.Rect{X: 50, Y: 50, W: w - 100, H: h - 100}

	if err := renderer.SetDrawColor(210, 180, 140, 255); err != nil {
		return err
	}

	return (renderer.FillRect(&r))
}

type card struct {
	rect  sdl.Rect
	color sdl.Color
	touch int64
}

var cards []card

func drawCards(renderer *sdl.Renderer) error {
	// Loop through cards in reverse, because we put the highest one at the front of the slice
	for i := range cards {
		c := cards[len(cards)-i-1]
		if err := renderer.SetDrawColorArray(c.color.R, c.color.G, c.color.B, c.color.A); err != nil {
			return err
		}
		if err := renderer.FillRect(&c.rect); err != nil {
			return err
		}
	}

	return nil
}

func checkCards(x, y int32) (*card, bool) {
	p := sdl.Point{X: x, Y: y}
	for i, c := range cards {
		if p.InRect(&c.rect) {
			return &cards[i], true
		}
	}
	return nil, false
}

var grabbedCard *card

var touch int64

func grabCard(c *card) {
	if grabbedCard != nil {
		panic(fmt.Sprintln("Already have card", *grabbedCard, "grabbed, can't grab another", *c))
	}
	log.Println("Grabbed card", *c)
	grabbedCard = c

	touch++
	c.touch = touch

	sortCards()

	grabbedCard = &cards[0]
}

func releaseCard() {
	if grabbedCard != nil {
		log.Println("Released card", *grabbedCard)
		grabbedCard = nil
	}
}

func dragCard(dX, dY int32) error {
	if grabbedCard == nil {
		return nil
	}

	grabbedCard.rect.X += dX
	grabbedCard.rect.Y += dY

	return nil
}

func sortCards() {
	sort.SliceStable(cards, func(i, j int) bool {
		return cards[i].touch > cards[j].touch
	})
}
