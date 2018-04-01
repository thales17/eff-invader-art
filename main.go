package main

import (
	"math/rand"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	windowW   = 1024
	windowH   = 768
	pixelSize = 2
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
		frames:     frames,
		frameIndex: 0,
		color:      color,
	}

	s.drawFrame()
	return s
}

func main() {
	canvas := sdl.NewCanvas("Eff Invader Art", windowW, windowH, eff.Color{R: 0x00, B: 0x00, G: 0x00, A: 0xFF}, 60, true)
	canvas.Run(func() {
		canvas.SetPrintFPS(true)
		rand.Seed(time.Now().UnixNano())
		white := eff.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}
		sizeX := 11 * pixelSize
		sizeY := 8 * pixelSize
		padding := 1 * pixelSize
		w := sizeX + (padding * 2)
		h := sizeY + (padding * 2)
		cols := (windowW / w) + 2
		rows := (windowH / h) + 2
		invaders := []*sprite{}
		for j := 0; j < rows; j++ {
			for i := 0; i < cols; i++ {
				invader := newSprite(Invader1, 10, white)
				invader.SetRect(eff.Rect{
					X: i * w,
					Y: j * h,
					W: sizeX,
					H: sizeY,
				})
				invaders = append(invaders, invader)
				canvas.AddChild(invader)
			}
		}

		ticks := 0
		frameTickCount := 10
		frameIndex := 0
		canvas.SetUpdateHandler(func() {
			ticks++
			if ticks == frameTickCount {
				ticks = 0
				frameIndex = (frameIndex + 1) % 2
				for _, inv := range invaders {
					inv.frameIndex = frameIndex
					inv.drawFrame()
				}
			}
		})
	})
}
