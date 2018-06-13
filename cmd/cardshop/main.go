package main

import (
	"log"

	"github.com/thomas-holmes/cardshop/pkg/game"
)

func main() {
	if err := game.Run(); err != nil {
		panic(err)
	}
	log.Println("Yo")
}
