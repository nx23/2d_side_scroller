package player

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nx23/2d_side_scroller/assets"
)

type Vector struct {
	X, Y float64
}

type Character struct {
	Position Vector
	W, H int
	speed     float64
	Sprites    map[string]*ebiten.Image
}

var Player = Character{
	Position: Vector{X: 100, Y: 100},
	W: 16,
	H: 16,
	speed: float64(300 / ebiten.TPS()),
	Sprites: map[string]*ebiten.Image{
		"body": assets.BodySprite,
		"face": assets.FaceSprite,
	},
}

func (player *Character) DrawPlayer(screen *ebiten.Image) {
	bodyOp := &ebiten.DrawImageOptions{}
	bodyOp.GeoM.Translate(player.Position.X, player.Position.Y)
	faceOp := &ebiten.DrawImageOptions{}
	faceOp.GeoM.Translate(player.Position.X+14, player.Position.Y+16)

	screen.DrawImage(player.Sprites["body"], bodyOp)
	screen.DrawImage(player.Sprites["face"], faceOp)
}

func (player *Character) Update() {
	var delta Vector

	if ebiten.IsKeyPressed(ebiten.KeyDown) && player.Position.Y < 400 {
		delta.Y = player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) && player.Position.Y > 0 {
		delta.Y = -player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && player.Position.X > 0 {
		delta.X = -player.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && player.Position.X < 560 {
		delta.X = player.speed
	}
	
	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := player.speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	player.Position.X += delta.X
	player.Position.Y += delta.Y
}