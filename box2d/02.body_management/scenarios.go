package main

var (
	world_paramCreateWorld = ParamCreateWorld{
		ID:       "world",
		GravityX: 0.0,
		GravityY: -10.0,
	}

	world_paramRunWorld = ParamRunWorld{
		VelocityIteration: 6,
		PositionIteration: 2,
		StepTime:          1.0 / 60.0,
		FrameLength:       uint(60),
	}

	groundBox_paramCreateBody = ParamCreateBody{
		Id:       "ground_box",
		PosX:     0.0,
		PosY:     -10.0,
		SizeX:    100.0,
		SizeY:    20.0,
		Density:  0.0,
		Friction: 0.0,
		Shape:    ShapeTypePolygon,
		Dynamic:  false,
	}

	simpleBox_paramCreateBody = ParamCreateBody{
		Id:       "simple_box",
		PosX:     0.0,
		PosY:     4.0,
		SizeX:    1.0,
		SizeY:    1.0,
		Density:  1.0,
		Friction: 0.3,
		Shape:    ShapeTypePolygon,
		Dynamic:  true,
	}
)

func simpleBox_falldown() {
	// scene
	scene := CreateScene("scene")

	// physics
	world := CreateWorld(world_paramCreateWorld)
	groundBox := CreateBody(groundBox_paramCreateBody)
	simpleBox := CreateBody(simpleBox_paramCreateBody)

	world.AddBody(groundBox)
	world.AddBody(simpleBox)
	world.Run(world_paramRunWorld, scene)

	// write output
	scene.WriteToFile("out.json")
}
