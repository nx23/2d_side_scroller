package sprite

import "github.com/hajimehoshi/ebiten/v2"

type Sprite struct {
	X, Y 	int
	Image *ebiten.Image
}
