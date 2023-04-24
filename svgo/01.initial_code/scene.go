package main

type Scene struct {
	Shapes []*Shape
	Frames []map[string]Pair3
}

func (s *Scene) LoadFrames() {
	s.Frames = []map[string]Pair3{}
	pos0 := map[string]Pair3{}
	for _, shape := range s.Shapes {
		pos0[shape.ID] = Pair3{
			X:     shape.Position.X,
			Y:     shape.Position.Y,
			Theta: shape.Position.Theta,
		}
		s.Frames = append(s.Frames, pos0)
	}
}

func (s *Scene) AddShape(shape *Shape) {
	s.Shapes = append(s.Shapes, shape)
}

func (s *Scene) Render(frameIndex int) {
	for _, shape := range s.Shapes {
		switch shape.Kind {
		case ShapeKindValuePolygon:
			if shape.Points == nil || len(shape.Points) == 0 {
				s.drawPolygonByPosAndSize(shape)
			} else {
				s.drawPolygonByPoints(shape)
			}
		case ShapeKindValueEllipse:
			s.drawEllipseByPosAndSize(shape)
		case ShapeKindValueNone:
		}
	}
}

func (s *Scene) drawPolygonByPosAndSize(shape *Shape) {

}

func (s *Scene) drawPolygonByPoints(shape *Shape) {

}

func (s *Scene) drawEllipseByPosAndSize(shape *Shape) {

}
