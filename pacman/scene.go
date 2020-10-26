package pacman

import (
	"github.com/hajimehoshi/ebiten"
)

type scene struct{}

func newScene() *scene {
	s := &scene{}
	return s
}

func (s *scene) update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebiten.DebugPrint(screen, "hola")
	return nil
}
