package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nx23/2d_side_scroller/assets"
)

type Vector struct {
	X, Y float64
}

type Character struct {
	Position Vector
	W, H int
	HorizontalSpeed     float64
	VerticalSpeed     float64
	Sprites    map[string]*ebiten.Image
}

var Player = Character{
	Position: Vector{X: 100, Y: 400},
	W: 16,
	H: 16,
	HorizontalSpeed: float64(300 / ebiten.TPS()),
	VerticalSpeed: 0,
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
	player.Move()
	player.Jump()
}

func (player *Character) Move() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) && player.Position.X > 0 {
		player.Position.X -= player.HorizontalSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) && player.Position.X < 560 {
		player.Position.X += player.HorizontalSpeed
	}
}

func (player *Character) Jump() {
	const jumpStrength = 10
	const gravity = 0.5
	const groundLevel = 400

	if ebiten.IsKeyPressed(ebiten.KeySpace) && player.Position.Y >= groundLevel {
		player.VerticalSpeed = -jumpStrength
	}

	player.Position.Y += player.VerticalSpeed
	if player.Position.Y < groundLevel {
		player.VerticalSpeed += gravity
	} else {
		player.Position.Y = groundLevel
		player.VerticalSpeed = 0
	}
}