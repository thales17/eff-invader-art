package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW   = 800
	windowH   = 540
	pixelSize = 4
)

type sprite struct {
	eff.Shape
	frames         [][]eff.Point
	frameIndex     int
	ticks          int
	frameTickCount int
	color          eff.Color
}

func (s *sprite) drawFrame() {
	s.Clear()
	for _, p := range s.frames[s.frameIndex] {
		s.FillRect(eff.Rect{
			X: p.X * pixelSize,
			Y: p.Y * pixelSize,
			W: pixelSize,
			H: pixelSize,
		}, s.color)
	}
}

func (s *sprite) tick() {
	s.ticks++
	if s.ticks == s.frameTickCount {
		s.ticks = 0
		s.frameIndex = (s.frameIndex + 1) % len(s.frames)
		s.drawFrame()
	}
}

func newSprite(frames [][]eff.Point, frameTickCount int, color eff.Color) *sprite {
	s := &sprite{
		frames:         frames,
		frameIndex:     0,
		ticks:          0,
		frameTickCount: frameTickCount,
		color:          color,
	}

	s.drawFrame()
	return s
}

func main() {
	canvas := sdl.NewCanvas("Eff Invader Art", windowW, windowH, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		white := eff.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
		invader := newSprite([][]eff.Point{
			{
				eff.Point{X: 0, Y: 6},
				eff.Point{X: 0, Y: 5},
				eff.Point{X: 0, Y: 4},
				eff.Point{X: 10, Y: 6},
				eff.Point{X: 10, Y: 5},
				eff.Point{X: 10, Y: 4},
				eff.Point{X: 1, Y: 3},
				eff.Point{X: 2, Y: 2},
				eff.Point{X: 9, Y: 3},
				eff.Point{X: 8, Y: 2},
				eff.Point{X: 3, Y: 2},
				eff.Point{X: 4, Y: 2},
				eff.Point{X: 5, Y: 2},
				eff.Point{X: 6, Y: 2},
				eff.Point{X: 7, Y: 2},
				eff.Point{X: 3, Y: 1},
				eff.Point{X: 2, Y: 0},
				eff.Point{X: 7, Y: 1},
				eff.Point{X: 8, Y: 0},
				eff.Point{X: 2, Y: 4},
				eff.Point{X: 2, Y: 5},
				eff.Point{X: 2, Y: 6},
				eff.Point{X: 1, Y: 4},
				eff.Point{X: 8, Y: 4},
				eff.Point{X: 8, Y: 5},
				eff.Point{X: 8, Y: 6},
				eff.Point{X: 9, Y: 4},
				eff.Point{X: 7, Y: 7},
				eff.Point{X: 6, Y: 7},
				eff.Point{X: 3, Y: 7},
				eff.Point{X: 4, Y: 7},
				eff.Point{X: 2, Y: 3},
				eff.Point{X: 8, Y: 3},
				eff.Point{X: 4, Y: 3},
				eff.Point{X: 5, Y: 3},
				eff.Point{X: 6, Y: 3},
				eff.Point{X: 3, Y: 5},
				eff.Point{X: 4, Y: 5},
				eff.Point{X: 5, Y: 5},
				eff.Point{X: 6, Y: 5},
				eff.Point{X: 7, Y: 5},
				eff.Point{X: 7, Y: 4},
				eff.Point{X: 6, Y: 4},
				eff.Point{X: 5, Y: 4},
				eff.Point{X: 4, Y: 4},
				eff.Point{X: 3, Y: 4},
			},
			{
				eff.Point{X: 1, Y: 7},
				eff.Point{X: 2, Y: 6},
				eff.Point{X: 9, Y: 7},
				eff.Point{X: 8, Y: 6},
				eff.Point{X: 9, Y: 5},
				eff.Point{X: 1, Y: 5},
				eff.Point{X: 2, Y: 5},
				eff.Point{X: 3, Y: 5},
				eff.Point{X: 4, Y: 5},
				eff.Point{X: 5, Y: 5},
				eff.Point{X: 6, Y: 5},
				eff.Point{X: 8, Y: 5},
				eff.Point{X: 7, Y: 5},
				eff.Point{X: 0, Y: 4},
				eff.Point{X: 10, Y: 4},
				eff.Point{X: 0, Y: 1},
				eff.Point{X: 0, Y: 2},
				eff.Point{X: 0, Y: 3},
				eff.Point{X: 10, Y: 3},
				eff.Point{X: 10, Y: 2},
				eff.Point{X: 10, Y: 1},
				eff.Point{X: 1, Y: 4},
				eff.Point{X: 1, Y: 3},
				eff.Point{X: 9, Y: 4},
				eff.Point{X: 9, Y: 3},
				eff.Point{X: 2, Y: 2},
				eff.Point{X: 8, Y: 2},
				eff.Point{X: 3, Y: 2},
				eff.Point{X: 4, Y: 2},
				eff.Point{X: 5, Y: 2},
				eff.Point{X: 6, Y: 2},
				eff.Point{X: 7, Y: 2},
				eff.Point{X: 3, Y: 1},
				eff.Point{X: 2, Y: 0},
				eff.Point{X: 7, Y: 1},
				eff.Point{X: 8, Y: 0},
				eff.Point{X: 2, Y: 3},
				eff.Point{X: 2, Y: 4},
				eff.Point{X: 3, Y: 4},
				eff.Point{X: 4, Y: 4},
				eff.Point{X: 5, Y: 4},
				eff.Point{X: 6, Y: 4},
				eff.Point{X: 7, Y: 4},
				eff.Point{X: 8, Y: 4},
				eff.Point{X: 8, Y: 3},
				eff.Point{X: 6, Y: 3},
				eff.Point{X: 4, Y: 3},
				eff.Point{X: 5, Y: 3},
			},
		}, 10, white)
		sizeX := 11 * pixelSize
		sizeY := 8 * pixelSize
		invader.SetRect(eff.Rect{
			X: (windowW - sizeX) / 2,
			Y: (windowH - sizeY) / 2,
			W: sizeX,
			H: sizeY,
		})
		minSpeed := 3
		maxSpeed := 10
		vec := eff.Point{X: rand.Intn(maxSpeed-minSpeed) + minSpeed, Y: rand.Intn(maxSpeed-minSpeed) + minSpeed}
		invader.SetUpdateHandler(func() {
			invader.tick()
			x := invader.Rect().X + vec.X
			y := invader.Rect().Y + vec.Y
			if x <= 0 || x >= (canvas.Rect().W-invader.Rect().W) {
				vec.X *= -1
			}

			if y <= 0 || y >= (canvas.Rect().H-invader.Rect().H) {
				vec.Y *= -1
			}

			invader.SetRect(eff.Rect{
				X: x,
				Y: y,
				W: invader.Rect().W,
				H: invader.Rect().H,
			})
		})
		canvas.AddChild(invader)
	})
}
