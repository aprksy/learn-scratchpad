package main

type PositionWithAngle struct {
	X     float64
	Y     float64
	Angle float64
}

/*
=====================================================================
*/

func CreateComponent(id string) *Component {
	return &Component{
		ID:        id,
		Positions: []*PositionWithAngle{},
	}
}

type Component struct {
	ID        string
	Positions []*PositionWithAngle
}

func (c *Component) AddPositionWithAngle(x, y, theta float64) {
	c.Positions = append(c.Positions, &PositionWithAngle{x, y, theta})
}

/*
=====================================================================
*/

func CreateScene(id string) *Scene {
	return &Scene{
		ID:         id,
		Components: []*Component{},
	}
}

type Scene struct {
	ID         string
	Components []*Component
}

func (s *Scene) AddComponent(id string) *Component {
	out := CreateComponent(id)
	s.Components = append(s.Components, out)
	return out
}
