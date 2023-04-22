package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type PositionWithAngle struct {
	X     float64
	Y     float64
	Angle float64
}

/*
=====================================================================
*/

type Frame map[string]*PositionWithAngle

func (f Frame) AddPositionWithAngle(id string, x, y, theta float64) {
	f[id] = &PositionWithAngle{x, y, theta}
}

/*
=====================================================================
*/

func CreateScene(id string) *Scene {
	return &Scene{
		ID:     id,
		Frames: []Frame{},
	}
}

type Scene struct {
	ID     string  `json:"id"`
	Frames []Frame `json:"frames"`
}

func (s *Scene) AddFrameForSingleObject(id string, x, y, theta float64) {
	frame := Frame{}
	frame.AddPositionWithAngle(id, x, y, theta)
	s.Frames = append(s.Frames, frame)
}

func (s *Scene) WriteToFile(fn string) {
	data, _ := json.Marshal(s)
	err := os.WriteFile(fn, data, 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
