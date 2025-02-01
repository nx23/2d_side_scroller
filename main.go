package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nx23/2d_side_scroller/internal/player"
	// "github.com/hajimehoshi/ebiten/v2/text/v2"
)

// var pressStart2P []byte
// var pressStart2PFaceSource *text.GoTextFaceSource

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
    Player *player.Character
}

func (g *Game) Update() error {
    g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ScoreFont := assets.ScoreFont

	// Player
    g.Player.DrawPlayer(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowTitle("Simple 2D Side Scroller")
	ebiten.SetWindowSize(screenWidth, screenHeight)
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(150, 200)
    player := player.Player

	g := &Game{
        Player: &player,
    }

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}