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
	pixelSize = 3
)

type sprite struct {
	frames         [][]eff.Point
	frameIndex     int
	ticks          int
	frameTickCount int
	point          eff.Point
}

func (s *sprite) getRects() []eff.Rect {
	rects := []eff.Rect{}
	for _, p := range s.frames[s.frameIndex] {
		rects = append(rects, eff.Rect{
			X: s.point.X + p.X*pixelSize,
			Y: s.point.Y + p.Y*pixelSize,
			W: pixelSize,
			H: pixelSize,
		})
	}

	return rects
}

func newSprite(frames [][]eff.Point, frameTickCount int) *sprite {
	s := &sprite{
		frames:     frames,
		frameIndex: 0,
	}
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
				invader := newSprite(Invader1, 10)
				invader.point = eff.Point{
					X: i * w,
					Y: j * h,
				}
				invaders = append(invaders, invader)
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
				canvas.Clear()
				rects := []eff.Rect{}
				for _, inv := range invaders {
					inv.frameIndex = frameIndex
					rects = append(rects, inv.getRects()...)
				}
				canvas.FillRects(rects, white)
			}
		})

		canvas.Clear()
		rects := []eff.Rect{}
		for _, inv := range invaders {
			rects = append(rects, inv.getRects()...)
		}
		canvas.FillRects(rects, white)
	})
}
