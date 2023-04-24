package main

import (
	"fmt"
	"os"

	svg "github.com/ajstarks/svgo/float"
)

func CreateScene(size, wsize, worig Pair2) *Scene {
	return &Scene{
		size:        size,
		worldOrigin: worig,
		worldSize:   wsize,
		ratio: Pair2{
			X: size.X / wsize.X, // pixels/meter
			Y: size.Y / wsize.Y, // pixels/meter
		},
		Shapes: []*Shape{},
		Frames: []map[string]Pair3{},
	}
}

type Scene struct {
	size        Pair2
	worldOrigin Pair2
	worldSize   Pair2
	ratio       Pair2
	canvas      *svg.SVG
	Shapes      []*Shape
	Frames      []map[string]Pair3
}

func (s *Scene) pair2WorldToScene(coord Pair2) Pair2 {
	return Pair2{
		X: s.worldOrigin.X + coord.X*s.ratio.X,
		Y: s.worldOrigin.Y + coord.Y*s.ratio.Y,
	}
}

func (s *Scene) pair3WorldToScene(coord Pair3) Pair3 {
	coord1 := Pair2{
		X: coord.X,
		Y: coord.Y,
	}
	out1 := s.pair2WorldToScene(coord1)
	return Pair3{
		X:     out1.X,
		Y:     out1.Y,
		Theta: coord.Theta,
	}
}

func (s *Scene) LoadFrames() {
	s.Frames = []map[string]Pair3{}
	pos0 := map[string]Pair3{}
	for _, shape := range s.Shapes {
		pos0[shape.ID] = s.pair3WorldToScene(shape.Position)
	}
	s.Frames = append(s.Frames, pos0)
}

func (s *Scene) AddShape(shape *Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) Render(frameIndex int) {
	file, err := os.Create(fmt.Sprintf("frame_%d.svg", frameIndex))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	s.canvas = svg.New(file)
	s.canvas.Start(s.size.X, s.size.Y)
	for _, shape := range s.Shapes {
		switch shape.Kind {
		case ShapeKindValuePolygon:
			if shape.Points == nil || len(shape.Points) == 0 {
				s.drawPolygon(shape, s.Frames[frameIndex][shape.ID])
			} else {
				s.drawPolygonByPoints(shape)
			}
		case ShapeKindValueEllipse:
			s.drawEllipseByPosAndSize(shape)
		case ShapeKindValueNone:
		}
	}
	s.canvas.End()
}

func (s *Scene) drawPolygon(shape *Shape, pos Pair3) {
	w := shape.Size.X
	h := shape.Size.Y
	s.canvas.TranslateRotate(pos.X, pos.Y, pos.Theta)
	s.canvas.CenterRect(0.0, 0.0, w, h, "fill:white;stroke:black")
	s.canvas.Gend()
}

func (s *Scene) drawPolygonByPoints(shape *Shape) {

}

func (s *Scene) drawEllipseByPosAndSize(shape *Shape) {

}
