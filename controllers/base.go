package controllers

import "math/rand"

type Event struct {
	Gid   int
	Wid   int
	Value interface{}
}

type FlightController interface {
	Init()
	Tick()
	Update(event Event)
	GetInitJSON() []byte
	GetTickJSON() []byte
}

func GetRandomController() FlightController {
	idx := rand.Int31n(5)
	idx = 4 // TODO remove this
	var c FlightController
	switch idx {
	case 0:
		c = new(ControlController)
	case 1:
		c = new(EECOMController)
	case 2:
		c = new(FIDOController)
	case 3:
		c = new(GNCController)
	case 4:
		c = new(SurgeonController)
	}
	c.Init()
	return c
}
