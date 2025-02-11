package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nx23/2d_side_scroller/internal/background"
	"github.com/nx23/2d_side_scroller/internal/character"
	"github.com/nx23/2d_side_scroller/internal/platform"
	"github.com/nx23/2d_side_scroller/internal/screen"

	// "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/nx23/2d_side_scroller/assets"
)

// var pressStart2P []byte
// var pressStart2PFaceSource *text.GoTextFaceSource

const JumpStrength = 10
const Gravity = 0.5

type Game struct {
    Player *character.Character
	Platform *platform.Platform
}

func (g *Game) Update() error {
	collideOnBottom := g.Platform.Hitbox.CollidesOnBottom(&g.Player.Hitbox)
	collideOnTop := g.Platform.Hitbox.CollidesOnTop(&g.Player.Hitbox)
	collideOnLeft := g.Platform.Hitbox.CollidesOnLeft(&g.Player.Hitbox)
	collideOnRight := g.Platform.Hitbox.CollidesOnRight(&g.Player.Hitbox)

	if collideOnTop {
		g.Player.Grounded = true
	}

    if !collideOnTop && g.Player.Y < 400 {
		g.Player.Grounded = false
	}

    g.Player.Update(collideOnBottom, collideOnTop, collideOnLeft, collideOnRight)
    g.Player.Jump(JumpStrength, Gravity)


	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ScoreFont := assets.ScoreFont

    // Background
    background.Draw(screen)

	// Player
    g.Player.Draw(screen)

	// Platform
	g.Platform.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screen.Width, screen.Height
}

func main() {
	ebiten.SetWindowTitle("Simple 2D Side Scroller")
	ebiten.SetWindowSize(screen.Width, screen.Height)

	player := character.Character{
		X: 100.0,
		Y: 400.0,
		Width: 80.0,
		Height: 80.0,
		HorizontalSpeed: float64(300.0 / ebiten.TPS()),
		VerticalSpeed: 0.0,
		Sprites: map[string]*ebiten.Image{
			"body": assets.BodySprite,
			"face": assets.FaceSprite,
		},
	}

	platform := platform.Platform{
		X: 400.0,
		Y: 400.0,
		Width: 200.0,
		Height: 30.0,
		HorizontalSpeed: float64(300.0 / ebiten.TPS()),
		VerticalSpeed: 0.0,
		Sprites: map[string]*ebiten.Image{},
	}

	g := &Game{
        Player: &player,
		Platform: &platform,
    }

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}