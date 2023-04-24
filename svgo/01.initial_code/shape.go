package main

import "encoding/json"

type ShapeKindValueType int

const (
	ShapeKindValueNone    ShapeKindValueType = 0
	ShapeKindValuePolygon ShapeKindValueType = 1
	ShapeKindValueEllipse ShapeKindValueType = 2

	ShapeKindCodeNone    string = "none"
	ShapeKindCodePolygon string = "polygon"
	ShapeKindCodeEllipse string = "ellipse"
)

var (
	ShapeKindValues = map[string]ShapeKindValueType{
		ShapeKindCodeNone:    ShapeKindValueNone,
		ShapeKindCodePolygon: ShapeKindValuePolygon,
		ShapeKindCodeEllipse: ShapeKindValueEllipse,
	}

	ShapeKindCodes = map[ShapeKindValueType]string{
		ShapeKindValueNone:    ShapeKindCodeNone,
		ShapeKindValuePolygon: ShapeKindCodePolygon,
		ShapeKindValueEllipse: ShapeKindCodeEllipse,
	}
)

type Shape struct {
	ID       string             `json:"id"`
	Kind     ShapeKindValueType `json:"-"`
	Position Pair3              `json:"position"`
	Size     Pair2              `json:"size"`
	Points   []Pair2            `json:"points"`
}

func (s *Shape) MarshalJSON() ([]byte, error) {
	type Alias Shape
	return json.Marshal(&struct {
		Kind string `json:"kind"`
		*Alias
	}{
		Kind:  ShapeKindCodes[s.Kind],
		Alias: (*Alias)(s),
	})
}

func (s *Shape) UnmarshalJSON(data []byte) error {
	type Alias Shape
	aux := &struct {
		Kind string `json:"kind"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}
	s.Kind = ShapeKindValues[aux.Kind]
	return nil
}
