package main

import (
	"github.com/ByteArena/box2d"
)

/*
=====================================================================
*/

type ShapeType uint

const (
	ShapeTypeNone    ShapeType = 0
	ShapeTypePolygon ShapeType = 1
	ShapeTypeCircle  ShapeType = 2
)

type ParamCreateBody struct {
	Id           string
	PosX, PosY   float64
	SizeX, SizeY float64
	Density      float64
	Friction     float64
	Shape        ShapeType
	Dynamic      bool
}

/*
---------------------------------------------------------------------
*/

func CreateBody(param ParamCreateBody) *Body {
	out := &Body{
		ID:      param.Id,
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
	VelocityIteration int
	PositionIteration int
	StepTime          float64
	FrameLength       uint
}

type ParamCreateWorld struct {
	ID       string
	GravityX float64
	GravityY float64
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
