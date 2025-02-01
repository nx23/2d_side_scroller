package background

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nx23/2d_side_scroller/assets"
)

func Draw(screen *ebiten.Image) {
	backgroundOp := &ebiten.DrawImageOptions{}
	backgroundOp.GeoM.Translate(0, 0)
	backgroundOp.ColorScale.Scale(1, 1, 1, 1)
	screen.Fill(color.White)
	screen.DrawImage(assets.BackgroundSprite, backgroundOp)
}