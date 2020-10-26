package pacman


import (
	"github.com/hajimehoshi/ebiten"
)


// Game contiene todos la data del juego
type Game struct {
	scene *pacman.scene
}

// NewGame es un constructor de game
func NewGame() *Game {
	g := &Game{}
	return g
}

//ScreenWidth devuelve la anchura de la pantalla
func (g *Game) ScreenWith() int {
	return 320
}

//ScreenHeight devuelve la altura de la pantalla
func (g *Game) ScreenHeight() int {
	return 240
}

// Update hace refresh a la pantalla
func (g *Game) Update(screen *ebiten.Image) error {
	if g.scene == nil {
		g.scene = pacman.newScene()
	}
	return g.scene.update(screen)
}
