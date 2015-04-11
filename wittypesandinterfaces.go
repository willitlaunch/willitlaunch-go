package main

type World struct {
  players map[string]Player
}

type Player struct {
  controller FlightController
}

type FlightController interface {
  init()
  getInitJSON() []byte
  tick()
  update(event Event)
  getTickJSON() []byte
}

type Event struct {
  gid int
  wid int
  value interface{}
}

type Game interface {
  init()
  update(event Event)
  tick()
  getInputsState() []interface{}
  getOutputsState() []interface{}
  getObjectives() []string
}

