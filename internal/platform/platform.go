package platform

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nx23/2d_side_scroller/internal/hitbox"
)

type Platform struct {
	X, Y, Width, Height, HorizontalSpeed, VerticalSpeed float64
	Hitbox hitbox.Hitbox
	Sprites map[string]*ebiten.Image
}

func (p *Platform) Draw(screen *ebiten.Image) {
	rectColor := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	x, y := float32(p.X), float32(p.Y)
	w, h := float32(p.Width), float32(p.Height)

	vector.DrawFilledRect(screen, x, y, w, h, rectColor, false)
}