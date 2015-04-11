package games

import "github.com/willitlaunch/willitlaunch-go"

type Game interface {
	init()
	tick()
	update(event Event)
	getInputsState() []interface{}
	getOutputsState() []interface{}
	getObjectives() []string
}
