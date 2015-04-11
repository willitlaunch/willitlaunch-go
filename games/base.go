package games

import "github.com/willitlaunch/willitlaunch-go"

type Game interface {
	Init()
	Tick()
	Update(event Event)
	GetInputsState() []interface{}
	GetOutputsState() []interface{}
	GetObjectives() []string
}
