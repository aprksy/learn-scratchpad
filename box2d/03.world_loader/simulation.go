package main

import (
	"encoding/json"

	"github.com/ByteArena/box2d"
)

/*
=====================================================================
*/

type ShapeType uint

const (
	ShapeTypeValueNone    ShapeType = 0
	ShapeTypeValuePolygon ShapeType = 1
	ShapeTypeValueCircle  ShapeType = 2

	ShapeTypeCodeNone    string = "none"
	ShapeTypeCodePolygon string = "polygon"
	ShapeTypeCodeCircle  string = "string"
)

var (
	ShapeTypeValues = map[string]ShapeType{
		ShapeTypeCodeNone:    ShapeTypeValueNone,
		ShapeTypeCodePolygon: ShapeTypeValuePolygon,
		ShapeTypeCodeCircle:  ShapeTypeValueCircle,
	}

	ShapeTypeCodes = map[ShapeType]string{
		ShapeTypeValueNone:    ShapeTypeCodeNone,
		ShapeTypeValuePolygon: ShapeTypeCodePolygon,
		ShapeTypeValueCircle:  ShapeTypeCodeCircle,
	}
)

type ParamCreateBody struct {
	ID       string    `json:"id"`
	PosX     float64   `json:"pos_x"`
	PosY     float64   `json:"pos_y"`
	SizeX    float64   `json:"size_x"`
	SizeY    float64   `json:"size_y"`
	Density  float64   `json:"density"`
	Friction float64   `json:"friction"`
	Shape    ShapeType `json:"shape_type"`
	Dynamic  bool      `json:"dynamic"`
}

func (p *ParamCreateBody) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	p.ID = v["id"].(string)
	p.PosX = v["pos_x"].(float64)
	p.PosY = v["pos_y"].(float64)
	p.SizeX = v["size_x"].(float64)
	p.SizeY = v["size_y"].(float64)
	p.Density = v["density"].(float64)
	p.Friction = v["friction"].(float64)
	p.Shape = ShapeTypeValues[v["shape_type"].(string)]
	p.Dynamic = v["dynamic"].(bool)

	return nil
}

func (p *ParamCreateBody) MarshalJSON() ([]byte, error) {
	type Alias ParamCreateBody
	return json.Marshal(&struct {
		Shape string `json:"shape_type"`
		*Alias
	}{
		Shape: ShapeTypeCodes[p.Shape],
		Alias: (*Alias)(p),
	})
}

/*
---------------------------------------------------------------------
*/

func CreateBody(param ParamCreateBody) *Body {
	out := &Body{
		ID:      param.ID,
		Dynamic: param.Dynamic,
	}

	out.bodyDef = box2d.MakeB2BodyDef()
	dynamicBox := box2d.MakeB2PolygonShape()
	dynamicBox.SetAsBox(param.SizeX/2, param.SizeY/2)

	out.fixtureDef = box2d.MakeB2FixtureDef()
	out.fixtureDef.Friction = param.Friction
	out.fixtureDef.Shape = &dynamicBox
	out.fixtureDef.Density = 0.0

	if param.Dynamic {
		out.bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
		out.fixtureDef.Density = param.Density
	}
	out.bodyDef.Position.Set(param.PosX, param.PosY)

	return out
}

type Body struct {
	ID         string
	bodyDef    box2d.B2BodyDef
	fixtureDef box2d.B2FixtureDef
	physBody   *box2d.B2Body
	Dynamic    bool
}

/*
=====================================================================
*/

type ParamRunWorld struct {
	VelocityIteration int     `json:"velocity_iteration"`
	PositionIteration int     `json:"position_iteration"`
	StepTime          float64 `json:"step_time"`
	FrameLength       uint    `json:"frame_length"`
}

type ParamCreateWorld struct {
	ID       string  `json:"id"`
	GravityX float64 `json:"gravity_x"`
	GravityY float64 `json:"gravity_y"`
}

/*
---------------------------------------------------------------------
*/

func CreateWorld(param ParamCreateWorld) *World {
	return &World{
		ID: param.ID,
		physWorld: box2d.MakeB2World(box2d.B2Vec2{
			X: param.GravityX,
			Y: param.GravityY,
		}),
		DynamicBodies: []*Body{},
		StaticBodies:  []*Body{},
	}
}

type World struct {
	ID            string
	physWorld     box2d.B2World
	DynamicBodies []*Body
	StaticBodies  []*Body
}

func (w *World) AddBody(body *Body) {
	b := w.physWorld.CreateBody(&body.bodyDef)
	b.CreateFixtureFromDef(&body.fixtureDef)
	body.physBody = b

	if body.Dynamic {
		w.DynamicBodies = append(w.DynamicBodies, body)
	} else {
		w.StaticBodies = append(w.StaticBodies, body)
	}
}

func (w *World) Run(params ParamRunWorld, scene *Scene) {
	for i := 0; i < int(params.FrameLength); i++ {
		w.physWorld.Step(params.StepTime, params.VelocityIteration, params.PositionIteration)
		for _, dynBody := range w.DynamicBodies {
			if dynBody.physBody == nil {
				continue
			}
			position := dynBody.physBody.GetPosition()
			angle := dynBody.physBody.GetAngle()
			if scene != nil {
				scene.AddFrameForSingleObject(dynBody.ID, position.X, position.Y, angle)
			}
		}
	}
}
