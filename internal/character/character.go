package character

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/nx23/2d_side_scroller/internal/hitbox"
)

type Character struct {
	X, Y, Width, Height float64
	Hitbox hitbox.Hitbox
	Grounded bool
	HorizontalSpeed     float64
	VerticalSpeed     float64
	Sprites    map[string]*ebiten.Image
}

func (character *Character) Draw(screen *ebiten.Image) {
	bodyOp := &ebiten.DrawImageOptions{}
	bodyOp.GeoM.Translate(character.X, character.Y)
	faceOp := &ebiten.DrawImageOptions{}
	faceOp.GeoM.Translate(character.X+14, character.Y+16)
	character.Hitbox = *hitbox.NewHitbox(float32(character.X), float32(character.Y), float32(character.Width), float32(character.Height))
	// character.Hitbox.Draw(screen)
	screen.DrawImage(character.Sprites["body"], bodyOp)
	screen.DrawImage(character.Sprites["face"], faceOp)
}

func (character *Character) Update(CollideOnBottom, CollideOnTop, CollideOnLeft, CollideOnRight bool) {
	character.Move(CollideOnBottom, CollideOnTop, CollideOnLeft, CollideOnRight)
	character.Hitbox.X = float32(character.X)
    character.Hitbox.Y = float32(character.Y)
}

func (character *Character) Move(CollideOnBottom, CollideOnTop, CollideOnLeft, CollideOnRight bool) {
	leftPressed := ebiten.IsKeyPressed(ebiten.KeyLeft)
	rightPressed := ebiten.IsKeyPressed(ebiten.KeyRight)

	canMoveLeft := character.X > 0
	noCollisionRight := !CollideOnRight
	collisionTop := CollideOnTop
	canMoveRight := character.X < 560
	noCollisionLeft := !CollideOnLeft

	if (leftPressed && canMoveLeft && noCollisionRight) || (leftPressed && canMoveLeft && collisionTop) {
		character.X -= character.HorizontalSpeed
	}

	if (rightPressed && canMoveRight && noCollisionLeft) || (rightPressed && canMoveRight && collisionTop) {
		character.X += character.HorizontalSpeed
	}
}

func (character *Character) Jump(JumpStrength, gravity float64) {

	if ebiten.IsKeyPressed(ebiten.KeySpace) && character.Grounded {
		character.VerticalSpeed = -JumpStrength
		character.Grounded = false
	}

	if !character.Grounded {
		character.Y += character.VerticalSpeed
		character.VerticalSpeed += gravity
	}

	if character.Y >= 400 {
		character.Y = 400
		character.Grounded = true
		character.VerticalSpeed = 0
	}
}