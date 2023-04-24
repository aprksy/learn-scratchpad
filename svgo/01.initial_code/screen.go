package main

type Screen struct {
	Size      Pair2
	ViewPorts []*ViewPort
}

type ViewPort struct {
	Offset Pair2
	Size   Pair2
	Ratio  float64 // pixels / meter
	Scene  *Scene
}
