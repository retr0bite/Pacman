package main

import (
	"github.com/hajimehoshi/ebiten"
	pacman2 "github.com/retr0bite/Pacman/pacman"
	"log"
)

func main() {
	g := pacman2.NewGame()
	if err := ebiten.Run(g.Update, g.ScreenWith(), g.ScreenHeight(), 2, "Pacman"); err != nil {
		log.Fatal(err)
	}
}
