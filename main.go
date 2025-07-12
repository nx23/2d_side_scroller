package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nx23/2d_side_scroller/assets"
	"github.com/nx23/2d_side_scroller/internal/background"
	"github.com/nx23/2d_side_scroller/internal/character"
	"github.com/nx23/2d_side_scroller/internal/hitbox"
	"github.com/nx23/2d_side_scroller/internal/platform"
	"github.com/nx23/2d_side_scroller/internal/screen"
)

const (
    JumpStrength = 10
    Gravity      = 0.5
    CoyoteTime   = 6
)

type Game struct {
    Player       *character.Character
    Platforms    []*platform.Platform
    CoyoteTimer  int
    DebugEnabled bool
    F1KeyPressed bool
    JumpKeyPressed bool
}

func (g *Game) handleInput() (bool, bool, bool, bool) {
    left := ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA)
    right := ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD)
    jump := ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp)
    
    f1CurrentlyPressed := ebiten.IsKeyPressed(ebiten.KeyF1)
    if f1CurrentlyPressed && !g.F1KeyPressed {
        g.DebugEnabled = !g.DebugEnabled
    }
    g.F1KeyPressed = f1CurrentlyPressed
    
    jumpJustPressed := jump && !g.JumpKeyPressed
    g.JumpKeyPressed = jump
    
    return left, right, jump, jumpJustPressed
}

func (g *Game) updatePlayerHitbox() {
    g.Player.Hitbox = *hitbox.NewHitbox(
        float32(g.Player.X), 
        float32(g.Player.Y), 
        float32(g.Player.Width), 
        float32(g.Player.Height),
    )
}

func (g *Game) handleCollisions() {
    newGroundedState := false
    
    for _, p := range g.Platforms {
        p.Hitbox = *hitbox.NewHitbox(float32(p.X), float32(p.Y), float32(p.Width), float32(p.Height))
        
        // Simple overlap check first to avoid expensive collision detection
        if g.Player.X >= p.X + p.Width || g.Player.X + g.Player.Width <= p.X ||
           g.Player.Y >= p.Y + p.Height || g.Player.Y + g.Player.Height <= p.Y {
            continue // No collision possible, skip this platform
        }
        
        collideOnBottom := p.Hitbox.CollidesOnBottom(&g.Player.Hitbox)
        collideOnTop := p.Hitbox.CollidesOnTop(&g.Player.Hitbox)
        collideOnLeft := p.Hitbox.CollidesOnLeft(&g.Player.Hitbox)
        collideOnRight := p.Hitbox.CollidesOnRight(&g.Player.Hitbox)

        // Calculate overlaps only once
        overlapTop := (g.Player.Y + g.Player.Height) - p.Y
        overlapBottom := (p.Y + p.Height) - g.Player.Y
        overlapLeft := (g.Player.X + g.Player.Width) - p.X
        overlapRight := (p.X + p.Width) - g.Player.X

        // Handle collisions with simplified logic
        if collideOnTop && g.Player.VerticalSpeed >= 0 && overlapTop <= overlapLeft && overlapTop <= overlapRight {
            newGroundedState = true
            g.Player.VerticalSpeed = 0
            g.Player.Y = float64(p.Y - g.Player.Height)
            g.updatePlayerHitbox()
        } else if collideOnBottom && g.Player.VerticalSpeed < 0 && overlapBottom <= overlapLeft && overlapBottom <= overlapRight {
            g.Player.VerticalSpeed = 0
            g.Player.Y = float64(p.Y + p.Height)
            g.updatePlayerHitbox()
        } else if collideOnLeft && overlapLeft <= overlapTop {
            g.Player.X = float64(p.X - g.Player.Width)
            g.updatePlayerHitbox()
        } else if collideOnRight && overlapRight <= overlapTop {
            g.Player.X = float64(p.X + p.Width)
            g.updatePlayerHitbox()
        }
    }
    
    g.Player.Grounded = newGroundedState
}

func (g *Game) handleCoyoteTime(wasGrounded bool) {
    if g.Player.Grounded {
        g.CoyoteTimer = 0
    } else if wasGrounded && !g.Player.Grounded {
        g.CoyoteTimer = CoyoteTime
    } else if g.CoyoteTimer > 0 {
        g.CoyoteTimer--
    }
}

func (g *Game) Update() error {
    left, right, _, jumpJustPressed := g.handleInput()
    wasGrounded := g.Player.Grounded
    
    g.Player.Jump(JumpStrength, Gravity)
    g.Player.Update(left, right, false, g.Player.Grounded)
    
    // Only update hitbox once per frame
    g.updatePlayerHitbox()
    
    g.handleCollisions()
    g.handleCoyoteTime(wasGrounded)
    
    if jumpJustPressed && (g.Player.Grounded || g.CoyoteTimer > 0) {
        g.Player.VerticalSpeed = -JumpStrength
        g.CoyoteTimer = 0
    }

    return nil
}

func (g *Game) getInputText() string {
    var keys []string
    if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
        keys = append(keys, "LEFT")
    }
    if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
        keys = append(keys, "RIGHT")
    }
    if ebiten.IsKeyPressed(ebiten.KeySpace) || ebiten.IsKeyPressed(ebiten.KeyUp) {
        keys = append(keys, "JUMP")
    }
    
    if len(keys) == 0 {
        return "NONE"
    }
    
    result := ""
    for i, key := range keys {
        if i > 0 {
            result += " "
        }
        result += key
    }
    return result
}

func (g *Game) getStateText() string {
    state := "AIRBORNE"
    if g.Player.Grounded {
        state = "GROUNDED"
    }
    if g.CoyoteTimer > 0 {
        state += " COYOTE"
    }
    return state
}

func (g *Game) drawDebugInfo(screen *ebiten.Image) {
    const (
        debugBgWidth  = 350
        debugBgHeight = 100
    )
    
    vector.DrawFilledRect(screen, 0, 0, debugBgWidth, debugBgHeight, color.RGBA{0, 0, 0, 180}, false)
    
    debugText := fmt.Sprintf("DEBUG\nKeys: %s\nState: %s\nPlayer Y: %.1f\nVertical Speed: %.1f", 
        g.getInputText(), g.getStateText(), g.Player.Y, g.Player.VerticalSpeed)
    
    ebitenutil.DebugPrintAt(screen, debugText, 10, 10)
}

func (g *Game) Draw(screen *ebiten.Image) {
    background.Draw(screen)
    
    for _, p := range g.Platforms {
        p.Draw(screen)
    }
    
    g.Player.Draw(screen)
    
    if g.DebugEnabled {
        g.drawDebugInfo(screen)
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screen.Width, screen.Height
}

func createPlatforms() []*platform.Platform {
    ground := &platform.Platform{
        X: 0.0, Y: 450.0, Width: 640.0, Height: 30.0,
        HorizontalSpeed: 0.0, VerticalSpeed: 0.0,
        Sprites: map[string]*ebiten.Image{},
    }

    platform1 := &platform.Platform{
        X: 200.0, Y: 400.0, Width: 200.0, Height: 30.0,
        HorizontalSpeed: float64(300.0 / ebiten.TPS()), VerticalSpeed: 0.0,
        Sprites: map[string]*ebiten.Image{},
    }

    platform2 := &platform.Platform{
        X: 400.0, Y: 350.0, Width: 200.0, Height: 30.0,
        HorizontalSpeed: float64(300.0 / ebiten.TPS()), VerticalSpeed: 0.0,
        Sprites: map[string]*ebiten.Image{},
    }

    return []*platform.Platform{ground, platform1, platform2}
}

func main() {
    ebiten.SetWindowTitle("Simple 2D Side Scroller")
    ebiten.SetWindowSize(screen.Width, screen.Height)

    player := character.Character{
        X: 50.0, Y: 350.0, Width: 80.0, Height: 80.0,
        HorizontalSpeed: float64(300.0 / ebiten.TPS()),
        VerticalSpeed: 0.0,
        Sprites: map[string]*ebiten.Image{
            "body": assets.BodySprite,
            "face": assets.FaceSprite,
        },
    }

    g := &Game{
        Player:    &player,
        Platforms: createPlatforms(),
    }

    if err := ebiten.RunGame(g); err != nil {
        panic(err)
    }
}