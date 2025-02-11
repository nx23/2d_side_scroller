package hitbox

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Hitbox struct {
    X, Y, Width, Height float32
}

func NewHitbox(x, y, width, height float32) *Hitbox {
    return &Hitbox{
        X:      x,
        Y:      y,
        Width:  width,
        Height: height,
    }
}

func (h *Hitbox) CollidesOnTop(other *Hitbox) bool {
	return h.Y < other.Y+other.Height &&
		h.Y > other.Y &&
		h.X < other.X+other.Width &&
		h.X+h.Width > other.X
}

func (h *Hitbox) CollidesOnBottom(other *Hitbox) bool {
	return h.Y+h.Height >= other.Y &&
		h.Y+h.Height <= other.Y+other.Height &&
		h.X <= other.X+other.Width &&
		h.X+h.Width >= other.X
}

func (h *Hitbox) CollidesOnLeft(other *Hitbox) bool {
	return h.X <= other.X + other.Width &&
		h.X >= other.X &&
		h.Y <= other.Y+other.Height &&
		h.Y+h.Height >= other.Y
}

func (h *Hitbox) CollidesOnRight(other *Hitbox) bool {
	return h.X+h.Width >= other.X &&
		h.X+h.Width <= other.X+other.Width &&
		h.Y <= other.Y+other.Height &&
		h.Y+h.Height >= other.Y
}

func (h *Hitbox) Draw(screen *ebiten.Image) {
	// Draw visualizes the hitbox for debugging purposes during development.
	// This method should be removed or disabled in production builds.
    vector.DrawFilledRect(screen, h.X-2, h.Y-2, h.Width+4, h.Height+4, color.RGBA{0xff, 0, 0, 0xff}, false)
}