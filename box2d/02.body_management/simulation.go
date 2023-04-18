package main

import (
	"time"

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
	if param.Dynamic {
		out.bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	}
	out.bodyDef.Position.Set(param.PosX, param.PosY)
	dynamicBox := box2d.MakeB2PolygonShape()
	dynamicBox.SetAsBox(param.SizeX/2, param.SizeY/2)

	out.fixtureDef = box2d.MakeB2FixtureDef()
	out.fixtureDef.Shape = &dynamicBox
	out.fixtureDef.Density = param.Density
	out.fixtureDef.Friction = param.Friction

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

type RunParams struct {
	VelocityIteration int
	PositionIteration int
	StepTime          float64
	Duration          time.Duration
	OutputFilename    string
}

/*
---------------------------------------------------------------------
*/

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

func (w *World) Run(params RunParams) *Scene {
	scene := CreateScene(w.ID)
	for i := time.Duration(0); i < params.Duration; i++ {
		w.physWorld.Step(params.StepTime, params.VelocityIteration, params.PositionIteration)
		for _, dynBody := range w.DynamicBodies {
			if dynBody.physBody == nil {
				continue
			}
			position := dynBody.physBody.GetPosition()
			angle := dynBody.physBody.GetAngle()
			component := scene.AddComponent(dynBody.ID)
			component.AddPositionWithAngle(position.X, position.Y, angle)
		}
	}
	return scene
}
