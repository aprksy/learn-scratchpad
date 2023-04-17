package main

import (
	"fmt"

	"github.com/ByteArena/box2d"
)

var ()

func run() {
	gravity := box2d.B2Vec2{
		X: 0.0,
		Y: -10.0,
	}
	world := box2d.MakeB2World(gravity)
	groundBodyDef := box2d.MakeB2BodyDef()
	groundBodyDef.Position.Set(0.0, -10.0)
	groundBody := world.CreateBody(&groundBodyDef)
	groundBox := box2d.MakeB2PolygonShape()
	groundBox.SetAsBox(50.0, 10.0)
	groundBody.CreateFixture(&groundBox, 0.0)

	bodyDef := box2d.MakeB2BodyDef()
	bodyDef.Type = box2d.B2BodyType.B2_dynamicBody
	bodyDef.Position.Set(0.0, 4.0)

	dynamicBox := box2d.MakeB2PolygonShape()
	dynamicBox.SetAsBox(1.0, 1.0)
	fixtureDef := box2d.MakeB2FixtureDef()
	fixtureDef.Shape = &dynamicBox
	fixtureDef.Density = 1.0
	fixtureDef.Friction = 0.3

	body := world.CreateBody(&bodyDef)
	body.CreateFixtureFromDef(&fixtureDef)

	velocityIterations := 6
	positionIterations := 2
	timeStep := 1.0 / 60.0

	for i := 0; i < 60; i++ {
		world.Step(timeStep, velocityIterations, positionIterations)
		position := body.GetPosition()
		angle := body.GetAngle()
		fmt.Printf("%4.2f %4.2f %4.2f\n", position.X, position.Y, angle)
	}
}
