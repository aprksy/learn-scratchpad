package main

import (
	"fmt"
	"time"

	"github.com/ByteArena/box2d"
	"github.com/tidwall/sjson"
)

type RunParams struct {
	VelocityIteration int
	PositionIteration int
	StepTime          float64
	Duration          time.Duration
	OutputFilename    string
}

type Body struct {
	Id       string
	physBody *box2d.B2Body
}

type World struct {
	Id            string
	physWorld     box2d.B2World
	DynamicBodies []*Body
}

func (w *World) Run(params RunParams) *string {
	out, _ := sjson.Set("", "world_id", w.Id)
	for i := time.Duration(0); i < params.Duration; i++ {
		w.physWorld.Step(params.StepTime, params.VelocityIteration, params.PositionIteration)
		for _, dynBody := range w.DynamicBodies {
			if dynBody.physBody == nil {
				continue
			}
			position := dynBody.physBody.GetPosition()
			angle := dynBody.physBody.GetAngle()
			data := map[string]float64{
				"x":     position.X,
				"y":     position.Y,
				"angle": angle,
			}
			sjson.Set(out, fmt.Sprintf("steps.-1.%s", dynBody.Id), data)
		}
	}
	return &out
}
