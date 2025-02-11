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
	p.Hitbox = *hitbox.NewHitbox(float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height))
	vector.DrawFilledRect(screen, float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height), color.RGBA{0xff, 0, 0, 0xff}, false)
	// platform.Hitbox.Draw(screen)
}