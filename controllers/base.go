package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/willitlaunch/willitlaunch-go/games"
	"github.com/willitlaunch/willitlaunch-go/widgets"
)

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
	GetIsReady() bool
	GetName() string
	SetName(s string)
	CheckObjectives() bool
}

func GetNextController(count int) FlightController {
	ctype := count % 5
	var c FlightController
	switch ctype {
	case 0:
		c = new(SurgeonController)
	case 1:
		c = new(EECOMController)
	case 2:
		c = new(FIDOController)
	case 3:
		c = new(GNCController)
	case 4:
		c = new(ControlController)
	}
	c.Init()
	c.SetName(fmt.Sprintf("%s %d", c.GetName(), count/5+1))
	return c
}

type FlightControllerBase struct {
	Name     string
	Games    []games.Game
	ReadyBtn widgets.Button
}

func (fc *FlightControllerBase) Init() {
	fc.ReadyBtn = widgets.Button{WidgetBase: widgets.WidgetBase{Gid: 999, Wid: 999, Label: "READY"}, Value: false}
}

func (fc *FlightControllerBase) Tick() {
	for _, game := range fc.Games {
		game.Tick()
	}
}

func (fc *FlightControllerBase) Update(event Event) {
	for _, game := range fc.Games {
		if game.GetGid() == event.Gid {
			evt := games.Event{Gid: event.Gid, Wid: event.Wid, Value: event.Value}
			game.Update(evt)
			return
		}
	}

	if event.Gid == 999 {
		fc.ReadyBtn.Value = event.Value.(bool)
	}
}

func (fc *FlightControllerBase) GetInitJSON() []byte {
	var inputStates []interface{}
	var outputStates []interface{}
	var objectives []interface{}
	for _, game := range fc.Games {
		for _, state := range game.GetOutputsState() {
			outputStates = append(outputStates, state)
		}
		for _, state := range game.GetInputsState() {
			inputStates = append(inputStates, state)
		}
		for _, objective := range game.GetObjectives() {
			objectives = append(objectives, objective)
		}
	}

	output := map[string]interface{}{
		"inputWidgets":  inputStates,
		"outputWidgets": outputStates,
		"objectives":    objectives,
		"name":          fc.Name,
	}
	out, err := json.Marshal(output)
	if err != nil {
		out = []byte{'E', 'r', 'r', 'o', 'r'}
		fmt.Println("error:", err)
		return out
	}
	return out
}

func (fc *FlightControllerBase) GetTickJSON() []byte {
	var outputStates []interface{}
	for _, game := range fc.Games {
		for _, state := range game.GetOutputsState() {
			outputStates = append(outputStates, state)
		}
	}
	output := map[string]interface{}{
		"outputWidgets": outputStates,
	}

	out, err := json.Marshal(output)
	if err != nil {
		out := []byte{'E', 'r', 'r', 'o', 'r'}
		fmt.Println("error:", err)
		return out
	}

	return out
}

func (fc *FlightControllerBase) CheckObjectives() bool {
	// Emptyset gives true!
	won := true
	for _, game := range fc.Games {
		won = game.CheckObjectives() && won
	}
	return won
}

func (fc *FlightControllerBase) GetIsReady() bool {
	return fc.ReadyBtn.Value
}

func (fc *FlightControllerBase) GetName() string {
	return fc.Name
}

func (fc *FlightControllerBase) SetName(s string) {
	fc.Name = s
}
